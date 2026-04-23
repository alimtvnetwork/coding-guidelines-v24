#!/usr/bin/env bash
# =====================================================================
# curl-shim.sh — drop-in replacement for `curl` used by the installer
# test harness. Reads a routing manifest from $SHIM_MANIFEST and serves
# canned responses based on URL match. Logs every request to
# $SHIM_LOG so tests can assert which URLs were touched (and which
# weren't, for the no-fallback case).
#
# Manifest format (one rule per line, first match wins):
#   <url-substring>\t<status>\t<body-file-or-empty>
# Lines beginning with '#' or blank lines are ignored.
#
# Status semantics:
#   200  → write body-file to -o path (or stdout if -o absent), exit 0
#   404  → exit 22 (curl's "HTTP error" exit when -f is set)
#   timeout → sleep 30s then exit 28 (used to verify deadlines)
#
# We translate only the flags the generated installers actually use:
#   -fsSL --retry N --retry-delay N -o PATH URL
#   HEAD-style (-I) requests are answered with status only, no body.
# =====================================================================
set -euo pipefail

: "${SHIM_MANIFEST:?SHIM_MANIFEST not set}"
: "${SHIM_LOG:?SHIM_LOG not set}"

OUT_PATH=""
URL=""
HEAD_ONLY=false

while [[ $# -gt 0 ]]; do
  case "$1" in
    -o) OUT_PATH="$2"; shift 2 ;;
    -I|--head) HEAD_ONLY=true; shift ;;
    --retry|--retry-delay|--retry-max-time|--max-time|--connect-timeout|-A|--user-agent|-H|--header)
      shift 2 ;;
    --silent|--show-error|--fail|--location|--insecure)
      shift ;;
    -*)
      # Bundled short flags like -fsSI / -fsSL / -sI. Detect 'I'
      # anywhere in the bundle to set HEAD_ONLY; otherwise just skip.
      [[ "$1" == *I* ]] && HEAD_ONLY=true
      shift ;;
    *)  URL="$1"; shift ;;
  esac
done

printf 'REQ\t%s\t%s\n' "$([[ "$HEAD_ONLY" == true ]] && echo HEAD || echo GET)" "$URL" >> "$SHIM_LOG"

match_rule() {
  # Echo "<status>\t<body-file>" of first matching rule, or empty.
  local line url_sub status body
  while IFS=$'\t' read -r url_sub status body; do
    [[ -z "$url_sub" || "$url_sub" == \#* ]] && continue
    if [[ "$URL" == *"$url_sub"* ]]; then
      printf '%s\t%s\n' "$status" "$body"
      return 0
    fi
  done < "$SHIM_MANIFEST"
}

rule="$(match_rule || true)"
if [[ -z "$rule" ]]; then
  printf 'RES\t404\t%s\n' "$URL" >> "$SHIM_LOG"
  exit 22
fi
status="${rule%%$'\t'*}"
body="${rule#*$'\t'}"

printf 'RES\t%s\t%s\n' "$status" "$URL" >> "$SHIM_LOG"

case "$status" in
  200)
    if [[ "$HEAD_ONLY" == true ]]; then
      exit 0
    fi
    if [[ -n "$OUT_PATH" ]]; then
      cp "$body" "$OUT_PATH"
    else
      cat "$body"
    fi
    exit 0
    ;;
  404)  exit 22 ;;
  timeout)
    sleep 30
    exit 28
    ;;
  *)    exit 1 ;;
esac