#!/usr/bin/env bash
# ci-runner.sh — Shared CLI wrapper for the 6 reusable CI guards.
#
# Dispatches to per-guard scripts using a uniform flag contract so that
# `check`, `lint`, and `test` phases share the same UX (exit codes,
# JSON output, baseline handling, verbosity) regardless of which
# underlying language tooling each guard wraps.
#
# See spec/12-cicd-pipeline-workflows/03-reusable-ci-guards/07-shared-cli-wrapper.md

set -uo pipefail

# ---------------------------------------------------------------------------
# Constants
# ---------------------------------------------------------------------------

readonly EXIT_OK=0
readonly EXIT_VIOLATION=1
readonly EXIT_TOOL_ERROR=2
readonly EXIT_USAGE_ERROR=64

readonly PHASE_CHECK="check"
readonly PHASE_LINT="lint"
readonly PHASE_TEST="test"
readonly PHASE_ALL="all"

readonly GUARDS_CHECK=("forbidden-names" "naming-baseline" "collisions")
readonly GUARDS_LINT=("lint-diff" "lint-suggest")
readonly GUARDS_TEST=("test-summary")

# ---------------------------------------------------------------------------
# Defaults (overridden by flags)
# ---------------------------------------------------------------------------

PHASE=""
GUARD=""
JSON_OUT=""
BASELINE_FILE=""
SOURCE_DIR=""
RESULTS_DIR=""
VERBOSE=0
FIX_MODE=0
SCRIPTS_DIR="${CI_RUNNER_SCRIPTS_DIR:-.github/scripts}"
CONFIG_FILE=""

# ---------------------------------------------------------------------------
# Logging helpers (single-purpose, < 8 lines each)
# ---------------------------------------------------------------------------

log_info() {
  if [ "$VERBOSE" -eq 1 ]; then
    printf '[ci-runner] %s\n' "$*" >&2
  fi
}

log_error() {
  printf '::error::[ci-runner] %s\n' "$*" >&2
}

log_section() {
  printf '\n=== %s ===\n' "$*" >&2
}

# ---------------------------------------------------------------------------
# Usage
# ---------------------------------------------------------------------------

print_usage() {
  cat "$(dirname "$0")/ci-runner-usage.txt"
}

# ---------------------------------------------------------------------------
# Flag parsing
# ---------------------------------------------------------------------------

is_known_phase() {
  case "$1" in
    "$PHASE_CHECK"|"$PHASE_LINT"|"$PHASE_TEST"|"$PHASE_ALL") return 0 ;;
    *) return 1 ;;
  esac
}

parse_flags() {
  while [ $# -gt 0 ]; do
    case "$1" in
      --phase)        PHASE="${2:-}";        shift 2 ;;
      --guard)        GUARD="${2:-}";        shift 2 ;;
      --config)       CONFIG_FILE="${2:-}";  shift 2 ;;
      --json)         JSON_OUT="${2:-}";     shift 2 ;;
      --baseline)     BASELINE_FILE="${2:-}"; shift 2 ;;
      --source-dir)   SOURCE_DIR="${2:-}";   shift 2 ;;
      --results-dir)  RESULTS_DIR="${2:-}";  shift 2 ;;
      --scripts-dir)  SCRIPTS_DIR="${2:-}";  shift 2 ;;
      --verbose)      VERBOSE=1;             shift ;;
      --fix)          FIX_MODE=1;            shift ;;
      --help|-h)      print_usage; exit "$EXIT_OK" ;;
      *)              log_error "unknown flag: $1"; print_usage; exit "$EXIT_USAGE_ERROR" ;;
    esac
  done
}

# Load config file (if provided) and apply its values as defaults.
# Explicit CLI flags take precedence over config-file values.
apply_config_file() {
  if [ -z "$CONFIG_FILE" ]; then
    return "$EXIT_OK"
  fi
  if [ ! -f "$CONFIG_FILE" ]; then
    log_error "config file not found: $CONFIG_FILE"
    return "$EXIT_TOOL_ERROR"
  fi
  log_info "loading config from $CONFIG_FILE"
  local env_lines
  if ! env_lines=$(node "$(dirname "$0")/ci-config.mjs" --config "$CONFIG_FILE" --emit env); then
    log_error "config loader failed for $CONFIG_FILE"
    return "$EXIT_TOOL_ERROR"
  fi
  # shellcheck disable=SC2086
  eval "$env_lines"
  # Apply config defaults only when CLI flag is empty.
  [ -z "$SCRIPTS_DIR" ] || [ "$SCRIPTS_DIR" = ".github/scripts" ] && \
    SCRIPTS_DIR="${CI_GUARDS_SCRIPTS_DIR:-$SCRIPTS_DIR}"
  return "$EXIT_OK"
}

validate_flags() {
  if [ -z "$PHASE" ]; then
    log_error "--phase is required"
    return "$EXIT_USAGE_ERROR"
  fi
  if ! is_known_phase "$PHASE"; then
    log_error "unknown phase: $PHASE (expected check|lint|test|all)"
    return "$EXIT_USAGE_ERROR"
  fi
  return "$EXIT_OK"
}

# ---------------------------------------------------------------------------
# Guard dispatchers — one function per pattern, < 15 lines each.
# Each returns 0 / 1 / 2 per the shared exit-code contract.
# ---------------------------------------------------------------------------

run_forbidden_names() {
  local script="$SCRIPTS_DIR/check-forbidden-names.sh"
  if [ ! -x "$script" ] && [ ! -f "$script" ]; then
    log_error "missing: $script"
    return "$EXIT_TOOL_ERROR"
  fi
  log_info "running forbidden-names guard against ${SOURCE_DIR:-<default>}"
  bash "$script" "${SOURCE_DIR:-}"
}

run_naming_baseline() {
  local script="$SCRIPTS_DIR/check-naming.sh"
  if [ ! -f "$script" ]; then
    log_error "missing: $script"
    return "$EXIT_TOOL_ERROR"
  fi
  if [ "$FIX_MODE" -eq 1 ]; then
    log_info "regenerating naming baseline"
    bash "$script" --regenerate-baseline
    return $?
  fi
  CONST_DIR="${SOURCE_DIR:-}" BASELINE_FILE="${BASELINE_FILE:-}" bash "$script"
}

run_collisions() {
  local script="$SCRIPTS_DIR/check-collisions.py"
  if [ ! -f "$script" ]; then
    log_error "missing: $script"
    return "$EXIT_TOOL_ERROR"
  fi
  log_info "running collision audit"
  python3 "$script"
}

run_lint_diff() {
  local script="$SCRIPTS_DIR/lint-diff.py"
  local current="${LINT_CURRENT:-current.json}"
  if [ ! -f "$script" ]; then
    log_error "missing: $script"
    return "$EXIT_TOOL_ERROR"
  fi
  python3 "$script" --current "$current" --baseline "${BASELINE_FILE:-}"
}

run_lint_suggest() {
  local script="$SCRIPTS_DIR/lint-suggest.py"
  local current="${LINT_CURRENT:-current.json}"
  if [ ! -f "$script" ]; then
    log_error "missing: $script"
    return "$EXIT_TOOL_ERROR"
  fi
  python3 "$script" --current "$current" --baseline "${BASELINE_FILE:-}" \
    --out "${GITHUB_STEP_SUMMARY:-/dev/stdout}"
}

run_test_summary() {
  local script="$SCRIPTS_DIR/test-summary.sh"
  if [ -z "$RESULTS_DIR" ]; then
    log_error "test phase requires --results-dir"
    return "$EXIT_USAGE_ERROR"
  fi
  if [ ! -f "$script" ]; then
    log_error "missing: $script"
    return "$EXIT_TOOL_ERROR"
  fi
  bash "$script" "$RESULTS_DIR"
}

# ---------------------------------------------------------------------------
# Phase orchestration
# ---------------------------------------------------------------------------

dispatch_guard() {
  case "$1" in
    forbidden-names)  run_forbidden_names ;;
    naming-baseline)  run_naming_baseline ;;
    collisions)       run_collisions ;;
    lint-diff)        run_lint_diff ;;
    lint-suggest)     run_lint_suggest ;;
    test-summary)     run_test_summary ;;
    *) log_error "unknown guard: $1"; return "$EXIT_USAGE_ERROR" ;;
  esac
}

guards_for_phase() {
  case "$1" in
    "$PHASE_CHECK") printf '%s\n' "${GUARDS_CHECK[@]}" ;;
    "$PHASE_LINT")  printf '%s\n' "${GUARDS_LINT[@]}" ;;
    "$PHASE_TEST")  printf '%s\n' "${GUARDS_TEST[@]}" ;;
    "$PHASE_ALL")
      printf '%s\n' "${GUARDS_CHECK[@]}" "${GUARDS_LINT[@]}" "${GUARDS_TEST[@]}"
      ;;
  esac
}

aggregate_exit() {
  # Highest exit code wins: tool error (2) > violation (1) > ok (0)
  local current="$1"
  local incoming="$2"
  if [ "$incoming" -gt "$current" ]; then
    printf '%s' "$incoming"
  else
    printf '%s' "$current"
  fi
}

emit_json_summary() {
  if [ -z "$JSON_OUT" ]; then
    return
  fi
  log_info "writing JSON summary to $JSON_OUT"
  printf '%s\n' "$1" > "$JSON_OUT"
}

run_phase() {
  local phase="$1"
  local overall="$EXIT_OK"
  local results_json="["
  local separator=""
  while IFS= read -r guard; do
    [ -z "$guard" ] && continue
    log_section "guard: $guard"
    dispatch_guard "$guard"
    local code=$?
    overall=$(aggregate_exit "$overall" "$code")
    results_json="${results_json}${separator}{\"guard\":\"${guard}\",\"exit\":${code}}"
    separator=","
  done < <(guards_for_phase "$phase")
  results_json="${results_json}]"
  emit_json_summary "{\"phase\":\"${phase}\",\"overall\":${overall},\"guards\":${results_json}}"
  return "$overall"
}

# ---------------------------------------------------------------------------
# Main
# ---------------------------------------------------------------------------

main() {
  parse_flags "$@"
  if ! validate_flags; then
    exit "$EXIT_USAGE_ERROR"
  fi
  if ! apply_config_file; then
    exit "$EXIT_TOOL_ERROR"
  fi
  if [ -n "$GUARD" ]; then
    log_section "single-guard mode: $GUARD"
    dispatch_guard "$GUARD"
    local code=$?
    emit_json_summary "{\"guard\":\"${GUARD}\",\"exit\":${code}}"
    exit "$code"
  fi
  run_phase "$PHASE"
  exit $?
}

main "$@"