#!/usr/bin/env bash
# Provider/auth helpers for visibility-change.sh

get_origin_url() {
  git remote get-url origin 2>/dev/null || true
}

# Sets PROVIDER ('github'|'gitlab') or empty
resolve_provider() {
  local url="$1"
  PROVIDER=""
  [ -n "$url" ] || return 0
  local re='^(https?://|ssh://[^@]+@|[^@]+@)([^/:]+)'
  if [[ ! "$url" =~ $re ]]; then return 0; fi
  local host="${BASH_REMATCH[2]}"
  if [ "$host" = "github.com" ] || [ "$host" = "ssh.github.com" ]; then
    PROVIDER="github"; return 0
  fi
  if [ "$host" = "gitlab.com" ]; then PROVIDER="gitlab"; return 0; fi
  local allow="${VISIBILITY_GITLAB_HOSTS:-}"
  local h
  IFS=',' read -ra ALLOWED <<<"$allow"
  for h in "${ALLOWED[@]}"; do
    h="${h// /}"
    if [ -n "$h" ] && [ "$h" = "$host" ]; then PROVIDER="gitlab"; return 0; fi
  done
}

# Sets SLUG to "owner/repo" or empty
resolve_owner_repo() {
  local url="$1"
  local trimmed="${url%/}"
  trimmed="${trimmed%.git}"
  SLUG=""
  local re_https='^https?://[^/]+/([^/]+)/([^/]+)$'
  local re_scp='^[^@]+@[^:]+:([^/]+)/([^/]+)$'
  local re_ssh='^ssh://[^@]+@[^/]+/([^/]+)/([^/]+)$'
  if [[ "$trimmed" =~ $re_https ]]; then SLUG="${BASH_REMATCH[1]}/${BASH_REMATCH[2]}"; return 0; fi
  if [[ "$trimmed" =~ $re_scp   ]]; then SLUG="${BASH_REMATCH[1]}/${BASH_REMATCH[2]}"; return 0; fi
  if [[ "$trimmed" =~ $re_ssh   ]]; then SLUG="${BASH_REMATCH[1]}/${BASH_REMATCH[2]}"; return 0; fi
}

is_cli_available() {
  command -v "$1" >/dev/null 2>&1
}

# Echoes 'public'|'private' or empty on failure
get_current_visibility() {
  local provider="$1" slug="$2"
  if [ "$provider" = "github" ]; then
    gh repo view "$slug" --json visibility -q .visibility 2>/dev/null | tr '[:upper:]' '[:lower:]'
    return 0
  fi
  glab repo view "$slug" -F json 2>/dev/null \
    | awk -F'"' '/"visibility"[[:space:]]*:/ {print tolower($4); exit}'
}
