#!/usr/bin/env bash
# Repo-identity helpers for fix-repo.sh

get_repo_root() {
  git rev-parse --show-toplevel 2>/dev/null
}

get_remote_url() {
  local url
  url="$(git remote get-url origin 2>/dev/null || true)"
  if [ -n "$url" ]; then echo "$url"; return 0; fi
  git remote -v 2>/dev/null | awk '$3=="(fetch)"{print $2; exit}'
}

# Sets globals: PARSED_HOST, PARSED_OWNER, PARSED_REPO
parse_remote_url() {
  local url="$1"
  local re_https='^https?://([^/:]+)(:[0-9]+)?/([^/]+)/([^/]+?)(\.git)?(/.*)?$'
  local re_scp='^git@([^:]+):([^/]+)/([^/]+?)(\.git)?$'
  local re_ssh='^ssh://git@([^/:]+)(:[0-9]+)?/([^/]+)/([^/]+?)(\.git)?$'
  if [[ "$url" =~ $re_https ]]; then
    PARSED_HOST="${BASH_REMATCH[1]}"; PARSED_OWNER="${BASH_REMATCH[3]}"; PARSED_REPO="${BASH_REMATCH[4]}"; return 0
  fi
  if [[ "$url" =~ $re_scp ]]; then
    PARSED_HOST="${BASH_REMATCH[1]}"; PARSED_OWNER="${BASH_REMATCH[2]}"; PARSED_REPO="${BASH_REMATCH[3]}"; return 0
  fi
  if [[ "$url" =~ $re_ssh ]]; then
    PARSED_HOST="${BASH_REMATCH[1]}"; PARSED_OWNER="${BASH_REMATCH[3]}"; PARSED_REPO="${BASH_REMATCH[4]}"; return 0
  fi
  return 1
}

# Sets globals: SPLIT_BASE, SPLIT_VERSION
split_repo_version() {
  local repo="$1"
  local re='^(.+)-v([0-9]+)$'
  if [[ "$repo" =~ $re ]]; then
    SPLIT_BASE="${BASH_REMATCH[1]}"; SPLIT_VERSION="${BASH_REMATCH[2]}"; return 0
  fi
  return 1
}
