#!/usr/bin/env bash
# =====================================================================
# check-fix-repo-arg-forwarding.sh — verify run.sh forwards every flag
# to fix-repo.sh exactly, in order, with no mutation/quoting loss.
#
# Strategy: shim fix-repo.sh in a temp copy of the runner so it just
# echoes "$@" with NUL separators, then compare against the input.
#
# Spec: spec/15-distribution-and-runner/02-runner-contract.md
# =====================================================================
set -uo pipefail

ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
TMP="$(mktemp -d)"
trap 'rm -rf "$TMP"' EXIT

cp "$ROOT/run.sh" "$TMP/run.sh"
chmod +x "$TMP/run.sh"

# Shim fix-repo.sh: print each arg on its own line, prefixed.
cat > "$TMP/fix-repo.sh" <<'SHIM'
#!/usr/bin/env bash
for a in "$@"; do printf 'ARG=%s\n' "$a"; done
SHIM
chmod +x "$TMP/fix-repo.sh"

RC=0

assert_forward() {
    local label="$1"; shift
    local expected=""
    for a in "$@"; do expected+="ARG=${a}"$'\n'; done
    local actual
    actual="$("$TMP/run.sh" fix-repo "$@" 2>/dev/null)"
    if [ "$actual" != "${expected%$'\n'}" ]; then
        echo "::error::[$label] forwarding mismatch" >&2
        echo "  expected: $(printf '%q ' "$@")" >&2
        echo "  actual:   $actual" >&2
        RC=1
    fi
}

assert_forward "no-args"
assert_forward "single-mode" --3
assert_forward "dry-verbose" --dry-run --verbose
assert_forward "all-flags" --all --dry-run --verbose
assert_forward "spaced-value" --foo "bar baz"

if [ "$RC" -eq 0 ]; then
    echo "✅ run.sh forwards every fix-repo flag verbatim"
fi
exit "$RC"
