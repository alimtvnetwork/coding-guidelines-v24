#!/usr/bin/env bash
# =====================================================================
# check-visibility-provider-detect.sh — Phase 5 (visibility-change)
#
# Sources scripts/visibility-change/provider.sh and exercises the pure
# helper functions resolve_provider / resolve_owner_repo across all
# remote URL formats listed in spec §3.
# =====================================================================
set -uo pipefail

ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
# shellcheck source=scripts/visibility-change/provider.sh
. "$ROOT/scripts/visibility-change/provider.sh"
RC=0

assert_pair() {
    local name="$1" url="$2" exp_provider="$3" exp_slug="$4" hosts_env="${5:-}"
    PROVIDER=""; SLUG=""
    VISIBILITY_GITLAB_HOSTS="$hosts_env" resolve_provider "$url"
    resolve_owner_repo "$url"
    if [ "$PROVIDER" != "$exp_provider" ]; then
        echo "::error::[$name] provider expected '$exp_provider', got '$PROVIDER' (url=$url)" >&2
        RC=1
    fi
    if [ "$SLUG" != "$exp_slug" ]; then
        echo "::error::[$name] slug expected '$exp_slug', got '$SLUG' (url=$url)" >&2
        RC=1
    fi
}

# GitHub — HTTPS, SSH-SCP, SSH-URL, with and without .git suffix
assert_pair "gh https"       "https://github.com/foo/bar.git"      "github" "foo/bar"
assert_pair "gh https plain" "https://github.com/foo/bar"          "github" "foo/bar"
assert_pair "gh scp"         "git@github.com:foo/bar.git"          "github" "foo/bar"
assert_pair "gh ssh url"     "ssh://git@github.com/foo/bar.git"    "github" "foo/bar"

# GitLab.com
assert_pair "gl https"       "https://gitlab.com/grp/proj.git"     "gitlab" "grp/proj"
assert_pair "gl scp"         "git@gitlab.com:grp/proj.git"         "gitlab" "grp/proj"

# Self-hosted GitLab requires explicit allow-list env var
assert_pair "gl self denied" "https://gl.example.com/g/p.git"      ""       "g/p"
assert_pair "gl self allowed" "https://gl.example.com/g/p.git"     "gitlab" "g/p" "gl.example.com"
assert_pair "gl self multi"  "https://git.acme.io/g/p.git"         "gitlab" "g/p" "gl.example.com,git.acme.io"

# Unsupported provider
assert_pair "bitbucket"      "https://bitbucket.org/foo/bar.git"   ""       "foo/bar"

# Garbage URL
PROVIDER="x"; SLUG="x"
resolve_provider ""
resolve_owner_repo ""
if [ -n "$PROVIDER" ] || [ -n "$SLUG" ]; then
    echo "::error::empty url should yield empty provider/slug, got '$PROVIDER' / '$SLUG'" >&2
    RC=1
fi

if [ "$RC" -eq 0 ]; then
    echo "✅ provider/slug resolution correct across 10 URL shapes"
fi
exit "$RC"
