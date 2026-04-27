#!/usr/bin/env bash
# =====================================================================
# check-fix-repo-url-rewrite.sh — locks the URL-embedded rewrite rule
#
# Spec:   spec-authoring/22-fix-repo/01-spec.md
# Memory: mem://features/fix-repo-url-handling
#         "fix-repo.ps1/.sh replace versioned-repo-name token in ALL
#          text files including inside URLs; host preserved automatically."
#
# Strategy: build a throw-away git repo whose origin is
# git@github.com:acme/widget-cli-v3.git, seed a README containing the
# token in 4 URL shapes + plain text, run fix-repo.sh --3, then assert:
#   - all `widget-cli-v1` and `widget-cli-v2` tokens became `widget-cli-v3`
#   - the surrounding URL hosts (github.com, gitlab.example.com,
#     raw.githubusercontent.com) are preserved verbatim
#   - `widget-cli-v3` (current) is NOT touched
#   - tokens followed by another digit (e.g. widget-cli-v10) are NOT
#     mistakenly rewritten (numeric overflow guard)
# =====================================================================
set -uo pipefail

ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
FR_SH="$ROOT/fix-repo.sh"
TMPDIR="$(mktemp -d)"
trap 'rm -rf "$TMPDIR"' EXIT
RC=0

if [ ! -f "$FR_SH" ]; then
    echo "::error::fix-repo.sh missing at $FR_SH" >&2
    exit 1
fi

cd "$TMPDIR"
git init -q -b main
git config user.email t@t.t
git config user.name t
git remote add origin "git@github.com:acme/widget-cli-v3.git"

cat > README.md <<'EOF'
# widget-cli-v3

Old GitHub HTTPS:   https://github.com/acme/widget-cli-v1
Old GitHub SSH-SCP: git@github.com:acme/widget-cli-v1.git
Old self-hosted:    https://gitlab.example.com/acme/widget-cli-v2/-/tree/main
Raw asset:          https://raw.githubusercontent.com/acme/widget-cli-v2/main/README.md
Plain reference:    See widget-cli-v1 docs.
Current (no-op):    widget-cli-v3 should stay.
Numeric guard:      widget-cli-v10 must NOT be rewritten.
EOF

hash="$(git hash-object -w README.md)"
git update-index --add --cacheinfo "100644,$hash,README.md"

bash "$FR_SH" --3 >/dev/null 2>&1
exit_code=$?
if [ "$exit_code" -ne 0 ]; then
    echo "::error::fix-repo.sh --3 exited $exit_code" >&2
    exit 1
fi

assert_grep() {
    local label="$1" pattern="$2"
    if ! grep -qF "$pattern" README.md; then
        echo "::error::[$label] expected to find: $pattern" >&2
        RC=1
    fi
}

assert_no_grep() {
    local label="$1" pattern="$2"
    if grep -qF "$pattern" README.md; then
        echo "::error::[$label] unexpected residual: $pattern" >&2
        RC=1
    fi
}

# --- Positive: rewritten URLs preserve their hosts -------------------
assert_grep "github https"   "https://github.com/acme/widget-cli-v3"
assert_grep "github ssh-scp" "git@github.com:acme/widget-cli-v3.git"
assert_grep "self-hosted"    "https://gitlab.example.com/acme/widget-cli-v3/-/tree/main"
assert_grep "raw asset"      "https://raw.githubusercontent.com/acme/widget-cli-v3/main/README.md"
assert_grep "plain text"     "See widget-cli-v3 docs."

# --- Negative: old tokens fully gone ---------------------------------
assert_no_grep "old v1 residue" "widget-cli-v1 "
assert_no_grep "old v1 url"     "widget-cli-v1.git"
assert_no_grep "old v2 residue" "widget-cli-v2/"

# --- Numeric guard: -v10 must survive untouched ----------------------
assert_grep "v10 preserved"  "widget-cli-v10 must NOT be rewritten."

if [ "$RC" -eq 0 ]; then
    echo "✅ fix-repo rewrites tokens inside URLs; hosts preserved; -v10 guarded"
fi
exit "$RC"
