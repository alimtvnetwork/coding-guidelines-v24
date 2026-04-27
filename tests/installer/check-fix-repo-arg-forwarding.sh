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

# Shim fix-repo.sh: emit ARGC then each arg on its own NUL-safe line.
# Using a sentinel <<EOA>> around each arg catches leading/trailing
# whitespace loss that a bare `ARG=…` line would hide.
cat > "$TMP/fix-repo.sh" <<'SHIM'
#!/usr/bin/env bash
printf 'ARGC=%d\n' "$#"
for a in "$@"; do printf 'ARG<<%s>>\n' "$a"; done
SHIM
chmod +x "$TMP/fix-repo.sh"

RC=0
PASS=0; FAIL=0

assert_forward() {
    local label="$1"; shift
    local expected="ARGC=$#"$'\n'
    for a in "$@"; do expected+="ARG<<${a}>>"$'\n'; done
    local actual
    actual="$("$TMP/run.sh" fix-repo "$@" 2>/dev/null)"
    if [ "$actual" = "${expected%$'\n'}" ]; then
        printf '  ✅ %s (argc=%d)\n' "$label" "$#"
        PASS=$((PASS + 1))
        return
    fi
    printf '  ❌ %s — forwarding mismatch\n' "$label" >&2
    printf '     input:    %s\n' "$(printf '%q ' "$@")" >&2
    printf '     expected: %q\n' "${expected%$'\n'}" >&2
    printf '     actual:   %q\n' "$actual" >&2
    FAIL=$((FAIL + 1))
    RC=1
}

# ── Baseline cases ─────────────────────────────────────────────────
assert_forward "no-args"
assert_forward "single-mode"          --3
assert_forward "dry-verbose"          --dry-run --verbose
assert_forward "all-flags"            --all --dry-run --verbose

# ── Tricky cases (the point of this PR) ───────────────────────────
# Single arg with embedded space — classic IFS-splitting trap.
assert_forward "value-with-spaces"    --foo "bar baz"

# --dry-run alongside a multi-word value (the user's example).
assert_forward "dryrun+spaces"        --dry-run --label "two words"

# Multiple words inside one arg + a separate flag after it.
assert_forward "spaces-then-flag"     "hello world" --verbose

# Glob characters MUST survive (no pathname expansion mid-flight).
assert_forward "glob-chars"           --pattern "*.md" --pattern "src/**/*.ts"

# Equals signs inside the value (looks like --flag=value but isn't).
assert_forward "equals-in-value"      --filter "key=value with spaces"

# Leading dashes inside a value (must not be re-parsed as flags).
assert_forward "value-looks-like-flag" --note "--this is not a flag"

# Tabs and newlines inside a single arg — the meanest IFS test.
assert_forward "tab-and-newline"      --raw $'col1\tcol2\nrow2'

# Single-quote and double-quote characters inside the value.
assert_forward "embedded-quotes"      --msg "she said \"hi\" and 'bye'"

# Empty string MUST be preserved as a distinct argv slot.
assert_forward "empty-string-arg"     --label "" --verbose

# A long mixed payload combining everything above.
assert_forward "mixed-payload" \
    --dry-run \
    --pattern "*.{md,ts}" \
    --note "value with  double  spaces" \
    --3

printf '\n  PASS=%d  FAIL=%d\n' "$PASS" "$FAIL"
if [ "$RC" -eq 0 ]; then
    echo "✅ run.sh forwards every fix-repo flag verbatim (incl. tricky argv)"
fi
exit "$RC"
