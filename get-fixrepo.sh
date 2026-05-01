#!/usr/bin/env bash
# ──────────────────────────────────────────────────────────────────────
# get-fixrepo.sh — installer for the fix-repo toolkit.
#
# Downloads from: github.com/alimtvnetwork/coding-guidelines-v20 (branch: main)
#
# Installs into the current working directory:
#   - fix-repo.sh, fix-repo.ps1
#   - fix-repo.config.json
#   - scripts/fix-repo/{repo-identity,file-scan,rewrite,config}.sh
#   - scripts/fix-repo/{RepoIdentity,FileScan,Rewrite,Config}.ps1
#   - fix-repo-contract.md  (the spec MD, dropped at root)
#
# After install, if CWD is inside a git repo, runs `./fix-repo.sh --dry-run`
# as a safe preview. If not a git repo, skips the run and exits 0.
#
# Usage:
#   curl -fsSL https://raw.githubusercontent.com/alimtvnetwork/coding-guidelines-v20/main/get-fixrepo.sh | bash
#   ./get-fixrepo.sh                      # install + dry-run preview
#   ./get-fixrepo.sh --no-run             # install only
#   ./get-fixrepo.sh --branch <name>      # install from a different branch
# ──────────────────────────────────────────────────────────────────────

set -euo pipefail

REPO="alimtvnetwork/coding-guidelines-v20"
BRANCH="main"
NO_RUN=0
RAW_BASE=""

EXIT_OK=0
EXIT_BAD_FLAG=6
EXIT_DOWNLOAD_FAILED=10
EXIT_NO_FETCHER=11

print_help() {
  sed -n '2,22p' "$0" | sed 's/^# \{0,1\}//'
}

parse_args() {
  while [ $# -gt 0 ]; do
    case "$1" in
      --no-run)   NO_RUN=1 ;;
      --branch)   shift; BRANCH="${1:-}" ;;
      --branch=*) BRANCH="${1#--branch=}" ;;
      -h|--help)  print_help; exit 0 ;;
      *) echo "get-fixrepo: ERROR unknown flag '$1'" >&2; exit $EXIT_BAD_FLAG ;;
    esac
    shift || true
  done
  [ -n "$BRANCH" ] || { echo "get-fixrepo: ERROR --branch requires a value" >&2; exit $EXIT_BAD_FLAG; }
  RAW_BASE="https://raw.githubusercontent.com/${REPO}/${BRANCH}"
}

has_curl() { command -v curl >/dev/null 2>&1; }
has_wget() { command -v wget >/dev/null 2>&1; }

assert_fetcher_available() {
  has_curl && return 0
  has_wget && return 0
  echo "get-fixrepo: ERROR neither curl nor wget is installed" >&2
  exit $EXIT_NO_FETCHER
}

fetch_one() {
  local rel="$1" dest="$2"
  mkdir -p "$(dirname "$dest")"
  if has_curl; then
    curl -fsSL "${RAW_BASE}/${rel}" -o "$dest" \
      || { echo "get-fixrepo: ERROR download failed for $rel" >&2; return 1; }
    return 0
  fi
  wget -q -O "$dest" "${RAW_BASE}/${rel}" \
    || { echo "get-fixrepo: ERROR download failed for $rel" >&2; return 1; }
}

# Files to install: each line is "<remote-rel-path>::<local-rel-path>".
manifest() {
  cat <<'EOF'
fix-repo.sh::fix-repo.sh
fix-repo.ps1::fix-repo.ps1
fix-repo.config.json::fix-repo.config.json
scripts/fix-repo/repo-identity.sh::scripts/fix-repo/repo-identity.sh
scripts/fix-repo/file-scan.sh::scripts/fix-repo/file-scan.sh
scripts/fix-repo/rewrite.sh::scripts/fix-repo/rewrite.sh
scripts/fix-repo/config.sh::scripts/fix-repo/config.sh
scripts/fix-repo/RepoIdentity.ps1::scripts/fix-repo/RepoIdentity.ps1
scripts/fix-repo/FileScan.ps1::scripts/fix-repo/FileScan.ps1
scripts/fix-repo/Rewrite.ps1::scripts/fix-repo/Rewrite.ps1
scripts/fix-repo/Config.ps1::scripts/fix-repo/Config.ps1
spec/02-coding-guidelines/06-cicd-integration/08-fix-repo-and-installers/01-fix-repo-contract.md::fix-repo-contract.md
EOF
}

install_all() {
  local count=0 line remote local
  while IFS= read -r line; do
    [ -n "$line" ] || continue
    remote="${line%%::*}"
    local="${line##*::}"
    fetch_one "$remote" "$local" || exit $EXIT_DOWNLOAD_FAILED
    count=$((count + 1))
  done < <(manifest)
  chmod +x fix-repo.sh 2>/dev/null || true
  echo "get-fixrepo: installed $count file(s) from ${REPO}@${BRANCH}"
}

is_git_repo() {
  git rev-parse --show-toplevel >/dev/null 2>&1
}

maybe_run_preview() {
  [ "$NO_RUN" = "0" ] || { echo "get-fixrepo: --no-run set; skipping preview"; return 0; }
  is_git_repo || {
    echo "get-fixrepo: not a git repository — files installed; skipping fix-repo run"
    return 0
  }
  echo
  echo "get-fixrepo: running './fix-repo.sh --dry-run' as a safe preview…"
  echo
  bash ./fix-repo.sh --dry-run || true
}

main() {
  parse_args "$@"
  assert_fetcher_available
  install_all
  maybe_run_preview
  exit $EXIT_OK
}

main "$@"
