#!/usr/bin/env bash
# ──────────────────────────────────────────────────────────────────────
# install.sh — Download spec/linters/scripts from a GitHub repo
#
# Conforms to: spec/14-update/27-generic-installer-behavior.md
#
# Quick start (defaults from install-config.json):
#   ./install.sh
#   curl -fsSL https://raw.githubusercontent.com/alimtvnetwork/coding-guidelines-v18/main/install.sh | bash
#
# Power-user flags:
#   --repo owner/repo            Override source repo
#   --branch main                Override branch (ignored if --version given)
#   --version vX.Y.Z             Install a specific release tag (PINNED MODE, §4)
#   --folders spec,linters       Comma-separated folder list (subpaths OK: spec/14-update)
#   --dest /path/to/dir          Install destination (default: cwd)
#   --log-dir /path or rel       Where to write fix-repo logs (default: <dest>/.install-logs;
#                                env: INSTALL_LOG_DIR)
#   --show-fix-repo-log          Print the latest fix-repo log to stdout after run_fix_repo
#                                completes (env: INSTALL_SHOW_FIX_REPO_LOG=1)
#   --config my-config.json      Use custom config file
#   --prompt                     Ask before overwriting each existing file (y/n/a/s)
#   --force                      Overwrite all existing files without prompting
#   --dry-run                    Show what would change; write nothing
#   --list-versions              List available release tags and exit
#   --list-folders               List available top-level folders for the chosen ref and exit
#   -n | --no-latest             Skip the latest-version probe (aliases: --no-probe)
#   --no-discovery               Skip V→V+N parallel discovery (spec §5.3)
#   --no-main-fallback           Skip main-branch fallback (spec §5.3)
#   --offline                    Skip all network ops; require local archive (alias: --use-local-archive)
#   --run-fix-repo               After verify, execute fix-repo.{sh,ps1} so the repo is patched
#                                before the installer exits (env: INSTALL_RUN_FIX_REPO=1)
#   -h | --help                  Show this help
#
# EXIT CODES (spec §8):
#   0  success
#   1  generic failure (missing tool, unknown flag, network exhausted)
#   2  offline mode required a network operation (or handshake mismatch)
#   3  pinned release / asset not found (PINNED MODE only)
#   4  verification failed (checksum / required-paths)
#   5  inner installer / handoff rejected
# ──────────────────────────────────────────────────────────────────────

set -euo pipefail

# ── Defaults ──────────────────────────────────────────────────────
CONFIG_FILE="install-config.json"
REPO=""
BRANCH=""
VERSION=""
DEST=""
FOLDERS=()
PROMPT_MODE=false
FORCE=false
DRY_RUN=false
LIST_VERSIONS=false
LIST_FOLDERS=false
PINNED_BY_RELEASE_INSTALL=""
NO_DISCOVERY=false
NO_MAIN_FALLBACK=false
OFFLINE=false
RUN_FIX_REPO="${INSTALL_RUN_FIX_REPO:-false}"
case "$RUN_FIX_REPO" in 1|true|TRUE|yes|YES) RUN_FIX_REPO=true ;; *) RUN_FIX_REPO=false ;; esac
ASSUME_YES="${INSTALL_FIX_REPO_YES:-false}"
case "$ASSUME_YES" in 1|true|TRUE|yes|YES) ASSUME_YES=true ;; *) ASSUME_YES=false ;; esac
ROLLBACK_ON_FIX_FAIL="${INSTALL_ROLLBACK_ON_FIX_REPO_FAILURE:-false}"
case "$ROLLBACK_ON_FIX_FAIL" in 1|true|TRUE|yes|YES) ROLLBACK_ON_FIX_FAIL=true ;; *) ROLLBACK_ON_FIX_FAIL=false ;; esac
FULL_ROLLBACK="${INSTALL_FULL_ROLLBACK:-false}"
case "$FULL_ROLLBACK" in 1|true|TRUE|yes|YES) FULL_ROLLBACK=true ;; *) FULL_ROLLBACK=false ;; esac
$FULL_ROLLBACK && ROLLBACK_ON_FIX_FAIL=true   # full implies edits
LOG_DIR="${INSTALL_LOG_DIR:-}"   # empty → $DEST/.install-logs (default)
SHOW_FIX_REPO_LOG="${INSTALL_SHOW_FIX_REPO_LOG:-false}"
case "$SHOW_FIX_REPO_LOG" in 1|true|TRUE|yes|YES) SHOW_FIX_REPO_LOG=true ;; *) SHOW_FIX_REPO_LOG=false ;; esac

# ── Colors ────────────────────────────────────────────────────────
RED='\033[0;31m'
GREEN='\033[0;32m'
CYAN='\033[0;36m'
YELLOW='\033[1;33m'
DIM='\033[2m'
NC='\033[0m'

step()  { echo -e "${CYAN}▸ $1${NC}"; }
ok()    { echo -e "${GREEN}✅ $1${NC}"; }
warn()  { echo -e "${YELLOW}⚠️  $1${NC}"; }
err()   { echo -e "${RED}❌ $1${NC}" >&2; }
dim()   { echo -e "${DIM}$1${NC}"; }

usage() {
  sed -n '2,26p' "$0" | sed 's/^# \{0,1\}//'
  exit 0
}

# ── Latest-version probe (unchanged) ──────────────────────────────
PROBE_OWNER_FALLBACK="alimtvnetwork"
PROBE_BASE_FALLBACK="coding-guidelines"
PROBE_VERSION_FALLBACK=14

invoke_latest_version_probe() {
    step "Detecting installer identity..."
    local src_url="${INSTALL_PROBE_SOURCE_URL:-${BASH_SOURCE[0]:-$0}}"
    local owner="${INSTALL_PROBE_OWNER:-}"
    local base="${INSTALL_PROBE_BASE:-}"
    local cur="${INSTALL_PROBE_VERSION:-}"
    local re='^https?://[^/]+/([^/]+)/([A-Za-z0-9._-]+)-v([0-9]+)/[^/]+/install\.sh'
    if [[ "$src_url" =~ $re ]]; then
        : "${owner:=${BASH_REMATCH[1]}}"
        : "${base:=${BASH_REMATCH[2]}}"
        : "${cur:=${BASH_REMATCH[3]}}"
    fi
    : "${owner:=$PROBE_OWNER_FALLBACK}"
    : "${base:=$PROBE_BASE_FALLBACK}"
    : "${cur:=$PROBE_VERSION_FALLBACK}"
    local current=$cur
    ok "Identity: $owner/$base-v$current  (probing v$((current+1))..v$((current+20)))"
    local depth=${INSTALL_PROBE_HANDOFF_DEPTH:-0}
    if [[ $depth -ge 3 ]]; then err "Probe loop guard (depth=$depth)"; exit 1; fi
    if ! command -v curl &>/dev/null; then warn "curl not found — skipping probe."; return 0; fi
    step "Probing 20 candidate versions in parallel (timeout 2s, middle-out)..."
    local tmp; tmp=$(mktemp -d)
    # Middle-out ordering: probe the middle of the range first, then expand
    # outward. With true parallelism this is correctness-equivalent, but it
    # plays better with any future early-abort logic (most active forks tend
    # to land mid-window of +1..+20).
    local low=$((current + 1))
    local high=$((current + 20))
    local mid=$(( (low + high) / 2 ))
    local order=("$mid")
    local offset upper lower
    for offset in $(seq 1 $((high - low))); do
        upper=$((mid + offset)); lower=$((mid - offset))
        [[ $upper -le $high ]] && order+=("$upper")
        [[ $lower -ge $low  ]] && order+=("$lower")
    done
    local n
    for n in "${order[@]}"; do
        (
            local url="https://raw.githubusercontent.com/$owner/$base-v$n/main/install.sh"
            local code; code=$(curl -s -o /dev/null -w '%{http_code}' --max-time 2 -I "$url" 2>/dev/null || echo 000)
            if [[ "$code" == "200" || "$code" == "301" || "$code" == "302" ]]; then echo "$n" > "$tmp/$n"; fi
        ) &
    done
    local waited=0
    while [[ $waited -lt 4 ]]; do sleep 1; waited=$((waited + 1)); [[ -z "$(jobs -rp)" ]] && break; done
    wait 2>/dev/null || true
    local latest=$current
    # Always pick the highest hit (sort -n | tail -1 → descending winner).
    if compgen -G "$tmp/*" >/dev/null 2>&1; then latest=$(basename "$(ls "$tmp" | sort -n | tail -1)"); fi
    rm -rf "$tmp"
    if [[ $latest -le $current ]]; then ok "Already on latest (v$current)."; return 0; fi
    local newer_url="https://raw.githubusercontent.com/$owner/$base-v$latest/main/install.sh"
    ok "Newer version found: v$latest. Handing off..."
    export INSTALL_PROBE_HANDOFF_DEPTH=$((depth + 1))
    export INSTALL_PROBE_SOURCE_URL="$newer_url"
    if curl -fsSL "$newer_url" | bash; then exit 0; else exit $?; fi
}

# Skip probe when user pinned a version, asked for a listing, or was
# launched by release-install.sh (handshake = --pinned-by-release-install).
should_skip_probe() {
  for arg in "$@"; do
    case "$arg" in
      --version|--list-versions|--list-folders|--no-probe|--no-latest|-n|--pinned-by-release-install|--run-fix-repo|-h|--help) return 0 ;;
    esac
  done
  [[ -n "${INSTALL_NO_PROBE:-}" ]] && return 0
  return 1
}

if ! should_skip_probe "$@"; then
  invoke_latest_version_probe || warn "Version probe error — continuing."
fi

# ── Parse CLI args ────────────────────────────────────────────────
while [[ $# -gt 0 ]]; do
  case "$1" in
    --repo)           REPO="$2";        shift 2 ;;
    --branch)         BRANCH="$2";      shift 2 ;;
    --version)        VERSION="$2";     shift 2 ;;
    --folders)        IFS=',' read -ra FOLDERS <<< "$2"; shift 2 ;;
    --dest)           DEST="$2";        shift 2 ;;
    --config)         CONFIG_FILE="$2"; shift 2 ;;
    --prompt)         PROMPT_MODE=true; shift ;;
    --force)          FORCE=true;       shift ;;
    --dry-run)        DRY_RUN=true;     shift ;;
    --list-versions)  LIST_VERSIONS=true; shift ;;
    --list-folders)   LIST_FOLDERS=true; shift ;;
    --no-probe|--no-latest|-n) shift ;;
    --no-discovery)   NO_DISCOVERY=true; shift ;;
    --no-main-fallback) NO_MAIN_FALLBACK=true; shift ;;
    --offline|--use-local-archive) OFFLINE=true; shift ;;
    --run-fix-repo)   RUN_FIX_REPO=true; shift ;;
    -y|--yes|--assume-yes) ASSUME_YES=true; shift ;;
    --rollback-on-fix-repo-failure) ROLLBACK_ON_FIX_FAIL=true; shift ;;
    --full-rollback)  FULL_ROLLBACK=true; ROLLBACK_ON_FIX_FAIL=true; shift ;;
    --log-dir)        LOG_DIR="$2"; shift 2 ;;
    --show-fix-repo-log) SHOW_FIX_REPO_LOG=true; shift ;;
    --pinned-by-release-install) PINNED_BY_RELEASE_INSTALL="$2"; shift 2 ;;
    -h|--help)        usage ;;
    *) err "Unknown option: $1"; exit 1 ;;
  esac
done

# Offline mode forbids any network operation (spec §5.3, §8 exit 2).
if $OFFLINE; then
  err "Offline mode is not yet supported by install.sh (no local-archive path). Exit 2 per spec §8."
  exit 2
fi

# Pinning handshake: when invoked by release-install.sh, the version
# arg MUST agree with the handshake value. Mismatch = exit 2 so the
# parent can detect version skew (per spec §Failure Modes).
if [[ -n "$PINNED_BY_RELEASE_INSTALL" ]]; then
  if [[ -z "$VERSION" ]]; then
    VERSION="$PINNED_BY_RELEASE_INSTALL"
  elif [[ "$VERSION" != "$PINNED_BY_RELEASE_INSTALL" ]]; then
    err "Pinning handshake mismatch: --version=$VERSION vs --pinned-by-release-install=$PINNED_BY_RELEASE_INSTALL"
    exit 2
  fi
  step "Pinned by release-install: $PINNED_BY_RELEASE_INSTALL (auto-update disabled)"
fi

if $PROMPT_MODE && $FORCE; then
  err "--prompt and --force are mutually exclusive"
  exit 1
fi

# ── Read config (only for repo/branch/folders defaults) ───────────
read_config() {
  local file="$1"
  [[ ! -f "$file" ]] && return 1
  local parser=""
  command -v python3 &>/dev/null && parser="python3"
  [[ -z "$parser" ]] && command -v node &>/dev/null && parser="node"
  if [[ -z "$parser" ]]; then err "Need python3 or node to parse $file"; exit 1; fi
  local result
  if [[ "$parser" == "python3" ]]; then
    result="$(python3 -c "import json; c=json.load(open('$file')); print(c.get('repo','')); print(c.get('branch','')); print('\n'.join(c.get('folders',[])))")"
  else
    result="$(node -e "const c=require('./$file'); console.log(c.repo||''); console.log(c.branch||''); (c.folders||[]).forEach(f=>console.log(f));")"
  fi
  local i=0
  while IFS= read -r line; do
    if [[ $i -eq 0 ]]; then [[ -z "$REPO" ]] && REPO="$line"
    elif [[ $i -eq 1 ]]; then [[ -z "$BRANCH" ]] && BRANCH="$line"
    else [[ ${#FOLDERS[@]} -eq 0 ]] && FOLDERS+=("$line")
    fi
    ((i++))
  done <<< "$result"
}

if [[ -f "$CONFIG_FILE" ]]; then
  step "Reading config from $CONFIG_FILE"
  read_config "$CONFIG_FILE"
fi

REPO="${REPO:-alimtvnetwork/coding-guidelines-v18}"
BRANCH="${BRANCH:-main}"
DEST="${DEST:-$(pwd)}"
[[ ${#FOLDERS[@]} -eq 0 ]] && FOLDERS=("spec" "linters" "linter-scripts" ".lovable/coding-guidelines")

# Top-level files always pulled alongside the folders. These are repo-root
# scripts (not contained in any installed folder) that users need locally to
# run repository hygiene tasks (fix-repo) and visibility toggles.
TOP_LEVEL_FILES=("fix-repo.sh" "fix-repo.ps1" "visibility-change.sh" "visibility-change.ps1")

# Ref = tag if --version, else branch
REF="$BRANCH"
[[ -n "$VERSION" ]] && REF="$VERSION"

# ── Download helpers ──────────────────────────────────────────────
download_to_file() {
  local url="$1" output="$2"
  if command -v curl &>/dev/null; then curl -fsSL "$url" -o "$output"
  elif command -v wget &>/dev/null; then wget -qO "$output" "$url"
  else err "Neither curl nor wget found"; exit 1; fi
}

download_to_stdout() {
  local url="$1"
  if command -v curl &>/dev/null; then curl -fsSL "$url"
  elif command -v wget &>/dev/null; then wget -qO- "$url"
  else err "Neither curl nor wget found"; exit 1; fi
}

# ── Listing modes (exit early) ────────────────────────────────────
list_release_versions() {
  step "Fetching releases for $REPO..."
  local json; json="$(download_to_stdout "https://api.github.com/repos/$REPO/releases?per_page=50" 2>/dev/null || echo "")"
  if [[ -z "$json" ]]; then err "Could not fetch releases"; exit 1; fi
  echo ""
  echo "$json" | grep -o '"tag_name"[[:space:]]*:[[:space:]]*"[^"]*"' | sed 's/.*"tag_name"[[:space:]]*:[[:space:]]*"//;s/"$//' | head -50 | while read -r tag; do
    echo "  • $tag"
  done
  echo ""
  exit 0
}

list_top_folders() {
  step "Listing folders for $REPO@$REF..."
  local api="https://api.github.com/repos/$REPO/contents?ref=$REF"
  local json; json="$(download_to_stdout "$api" 2>/dev/null || echo "")"
  if [[ -z "$json" ]]; then err "Could not list folders"; exit 1; fi
  echo ""
  echo "$json" | python3 -c "
import json, sys
items = json.load(sys.stdin)
dirs = sorted([i['name'] for i in items if i.get('type') == 'dir'])
for d in dirs: print(f'  • {d}')
" 2>/dev/null || echo "$json" | grep -B1 '"type": *"dir"' | grep '"name"' | sed 's/.*"name"[[:space:]]*:[[:space:]]*"//;s/".*$//' | sed 's/^/  • /'
  echo ""
  exit 0
}

$LIST_VERSIONS && list_release_versions
$LIST_FOLDERS && list_top_folders

# ── Banner (spec §7) ──────────────────────────────────────────────
INSTALL_MODE="implicit"
[[ -n "$VERSION" ]] && INSTALL_MODE="pinned"
SOURCE_KIND="tag-tarball"
[[ -z "$VERSION" ]] && SOURCE_KIND="branch-tarball"
echo ""
echo "    📦 Spec & Scripts Installer"
echo "       mode:    $INSTALL_MODE"
echo "       repo:    $REPO"
echo "       version: ${VERSION:-$BRANCH (implicit)}"
echo "       source:  $SOURCE_KIND"
echo "       folders: ${FOLDERS[*]}"
echo "       dest:    $DEST"
$DRY_RUN     && echo "       opts:    DRY-RUN (no writes)"
$PROMPT_MODE && echo "       opts:    Interactive prompts (y/n/a/s)"
$FORCE       && echo "       opts:    Force overwrite"
$NO_DISCOVERY && echo "       opts:    --no-discovery (V→V+N forbidden)"
$NO_MAIN_FALLBACK && echo "       opts:    --no-main-fallback"
echo ""

# ── Cleanup trap (with verification) ──────────────────────────────
TMP_DIR=""
cleanup() {
  [[ -z "${TMP_DIR:-}" ]] && return 0
  rm -rf "$TMP_DIR" 2>/dev/null || true
  if [[ ! -e "$TMP_DIR" ]]; then
    ok "Temp cleaned: $TMP_DIR"
  else
    warn "Temp NOT fully removed: $TMP_DIR (manual cleanup recommended)"
  fi
}
trap cleanup EXIT

# ── Fetch archive at REF (tarball, smaller) ───────────────────────
TMP_DIR="$(mktemp -d)"
ARCHIVE_PATH="$TMP_DIR/repo.tar.gz"
ARCHIVE_URL="https://codeload.github.com/$REPO/tar.gz/refs/heads/$BRANCH"
[[ -n "$VERSION" ]] && ARCHIVE_URL="https://codeload.github.com/$REPO/tar.gz/refs/tags/$VERSION"

step "Downloading $REPO@$REF..."
if ! download_to_file "$ARCHIVE_URL" "$ARCHIVE_PATH" 2>/dev/null; then
  # Fallback for tags that aren't refs/tags (some pipelines push as branch-like)
  ARCHIVE_URL="https://codeload.github.com/$REPO/tar.gz/$REF"
  download_to_file "$ARCHIVE_URL" "$ARCHIVE_PATH"
fi

step "Extracting..."
EXTRACT_DIR="$TMP_DIR/extracted"
mkdir -p "$EXTRACT_DIR"
tar -xzf "$ARCHIVE_PATH" -C "$EXTRACT_DIR"
ARCHIVE_ROOT="$(find "$EXTRACT_DIR" -mindepth 1 -maxdepth 1 -type d | head -1)"
[[ -z "$ARCHIVE_ROOT" ]] && { err "Failed to find archive root"; exit 1; }

# ── Merge with prompt/force/dry-run semantics ─────────────────────
PROMPT_ALL=false
PROMPT_SKIP_ALL=false
COPIED=0
SKIPPED_FILES=0
OVERWROTE=0
WROTE_NEW=0
SKIPPED_FOLDERS=0

# ── Rollback bookkeeping ──────────────────────────────────────────
# Populated by merge_file when not in --dry-run. Used by full rollback.
ROLLBACK_DIR=""           # set lazily by ensure_rollback_dir
INSTALLED_NEW=()          # paths created by this run (full path)
INSTALLED_BACKUPS=()      # "target|backup" pairs for overwritten files

ensure_rollback_dir() {
  [[ -n "$ROLLBACK_DIR" ]] && return 0
  local ts; ts="$(date -u +%Y%m%dT%H%M%SZ)"
  ROLLBACK_DIR="$DEST/.install-rollback/$ts"
  mkdir -p "$ROLLBACK_DIR/backups"
}

record_new_path() {
  ensure_rollback_dir
  INSTALLED_NEW+=("$1")
  echo "$1" >> "$ROLLBACK_DIR/new-paths.txt"
}

record_overwrite_backup() {
  local target="$1" rel
  ensure_rollback_dir
  rel="${target#$DEST/}"
  local backup="$ROLLBACK_DIR/backups/$rel"
  mkdir -p "$(dirname "$backup")"
  cp -f "$target" "$backup"
  INSTALLED_BACKUPS+=("$target|$backup")
  printf '%s\t%s\n' "$target" "$backup" >> "$ROLLBACK_DIR/backups.tsv"
}

decide_overwrite() {
  # Returns 0 to write, 1 to skip
  local target="$1"
  $PROMPT_ALL && return 0
  $PROMPT_SKIP_ALL && return 1
  $FORCE && return 0
  if ! $PROMPT_MODE; then return 0; fi  # default = overwrite (legacy behavior)
  while true; do
    echo -ne "${YELLOW}?${NC} Overwrite ${target#$DEST/} ? [y]es/[n]o/[a]ll/[s]kip-all: "
    read -r ans </dev/tty || ans="n"
    case "$ans" in
      y|Y) return 0 ;;
      n|N) return 1 ;;
      a|A) PROMPT_ALL=true; return 0 ;;
      s|S) PROMPT_SKIP_ALL=true; return 1 ;;
      *) echo "  enter y, n, a, or s" ;;
    esac
  done
}

merge_file() {
  local src="$1" target="$2"
  local target_dir; target_dir="$(dirname "$target")"
  if [[ -e "$target" ]]; then
    if decide_overwrite "$target"; then
      if $DRY_RUN; then
        dim "  ~ would overwrite ${target#$DEST/}"
      else
        $FULL_ROLLBACK && record_overwrite_backup "$target"
        mkdir -p "$target_dir"; cp -f "$src" "$target"
      fi
      ((OVERWROTE++))
    else
      dim "  - skip ${target#$DEST/}"
      ((SKIPPED_FILES++))
    fi
  else
    if $DRY_RUN; then
      dim "  + would create ${target#$DEST/}"
    else
      mkdir -p "$target_dir"; cp -f "$src" "$target"
      $FULL_ROLLBACK && record_new_path "$target"
    fi
    ((WROTE_NEW++))
  fi
}

merge_folder() {
  local folder="$1"
  local src="$ARCHIVE_ROOT/$folder"
  if [[ ! -d "$src" ]]; then
    warn "Folder '$folder' not found in $REPO@$REF — skipping"
    ((SKIPPED_FOLDERS++))
    return
  fi
  step "Merging: $folder"
  while IFS= read -r -d '' f; do
    local rel="${f#$src/}"
    merge_file "$f" "$DEST/$folder/$rel"
  done < <(find "$src" -type f -print0)
  ((COPIED++))
}

for folder in "${FOLDERS[@]}"; do
  merge_folder "$folder"
done

# Top-level files: copy each from the archive root into DEST (same name).
# Missing files are warned (not fatal) so installer remains forward-compatible
# with repos that omit a script.
for tlf in "${TOP_LEVEL_FILES[@]}"; do
  src="$ARCHIVE_ROOT/$tlf"
  if [[ ! -f "$src" ]]; then
    warn "Top-level file '$tlf' not found in $REPO@$REF — skipping"
    continue
  fi
  step "Merging file: $tlf"
  merge_file "$src" "$DEST/$tlf"
done

# ── Verify required files (spec §8: exit 4 on missing required path) ──
# Required files MUST exist in DEST after install. Skipped under --dry-run
# since no files were actually written.
REQUIRED_FILES=("fix-repo.sh" "fix-repo.ps1")
verify_required_files() {
  local missing=()
  local f
  for f in "${REQUIRED_FILES[@]}"; do
    [[ -f "$DEST/$f" ]] || missing+=("$f")
  done
  [[ ${#missing[@]} -eq 0 ]] && { ok "Verified ${#REQUIRED_FILES[@]} required file(s) present"; return 0; }
  err "Install verification FAILED — ${#missing[@]} required file(s) missing in $DEST:"
  local m
  for m in "${missing[@]}"; do echo "     • $m" >&2; done
  echo "" >&2
  echo "   The archive was downloaded but did NOT contain the expected" >&2
  echo "   top-level scripts. Re-run without --version to fetch main, or" >&2
  echo "   pin to a release that includes fix-repo.{sh,ps1}." >&2
  exit 4
}
$DRY_RUN || verify_required_files

# ── Optional: auto-run fix-repo so the repo is patched before exit ──
# Gated by --run-fix-repo or INSTALL_RUN_FIX_REPO=1. Picks .ps1 on
# Windows shells, .sh elsewhere. Skipped under --dry-run (nothing was
# actually written). Failures propagate as exit 5 per spec §8.
confirm_fix_repo() {
  $ASSUME_YES && { dim "Auto-confirmed (--yes / INSTALL_FIX_REPO_YES=1)"; return 0; }
  if [[ ! -t 0 ]]; then
    err "--run-fix-repo requires confirmation but stdin is not a TTY."
    err "   Re-run with --yes (or INSTALL_FIX_REPO_YES=1) to bypass the prompt."
    exit 5
  fi
  local reply=""
  warn "About to run $1"
  warn "This will rewrite versioned-repo-name tokens across tracked text files in this repo."
  printf "Proceed? [y/N] " >&2
  IFS= read -r reply </dev/tty || reply=""
  case "$reply" in y|Y|yes|YES) return 0 ;; esac
  warn "fix-repo skipped by user — exiting with code 5."
  exit 5
}
snapshot_pre_fix_repo() {
  # Capture HEAD ref of the destination repo so we can restore tracked
  # files via `git checkout HEAD -- .` if fix-repo fails.
  PRE_FIX_REPO_HEAD=""
  $ROLLBACK_ON_FIX_FAIL || return 0
  if ! git -C "$DEST" rev-parse --git-dir >/dev/null 2>&1; then
    warn "--rollback-on-fix-repo-failure: $DEST is not a git repo; rollback disabled."
    ROLLBACK_ON_FIX_FAIL=false
    return 0
  fi
  PRE_FIX_REPO_HEAD="$(git -C "$DEST" rev-parse HEAD 2>/dev/null || true)"
  step "Rollback armed: HEAD=$PRE_FIX_REPO_HEAD${FULL_ROLLBACK:+, full-rollback=on}"
}

restore_fix_repo_edits() {
  step "Restoring tracked files from HEAD..."
  git -C "$DEST" checkout -- . 2>&1 | tee -a "$1" || warn "git checkout reported issues"
  ok "fix-repo edits reverted"
}

restore_installed_paths() {
  local target backup pair removed=0 restored=0
  step "Removing files created by this install run..."
  for target in "${INSTALLED_NEW[@]:-}"; do
    [[ -e "$target" ]] || continue
    rm -f "$target" && ((removed++)) || warn "could not remove $target"
  done
  step "Restoring overwritten files from backups..."
  for pair in "${INSTALLED_BACKUPS[@]:-}"; do
    target="${pair%%|*}"; backup="${pair##*|}"
    [[ -f "$backup" ]] || continue
    cp -f "$backup" "$target" && ((restored++)) || warn "could not restore $target"
  done
  ok "Removed $removed new file(s); restored $restored overwritten file(s)"
}

perform_rollback() {
  local log="$1"
  echo ""
  warn "═══ ROLLBACK TRIGGERED (fix-repo failed) ═══"
  restore_fix_repo_edits "$log"
  $FULL_ROLLBACK && restore_installed_paths
  warn "Rollback complete. Snapshot kept at: ${ROLLBACK_DIR:-<none>}"
}

run_fix_repo() {
  local script log_dir log_file ts rc
  case "$(uname -s 2>/dev/null || echo unknown)" in
    MINGW*|MSYS*|CYGWIN*) script="$DEST/fix-repo.ps1" ;;
    *)                    script="$DEST/fix-repo.sh"  ;;
  esac
  if [[ ! -f "$script" ]]; then
    err "--run-fix-repo: $script not found after install."
    exit 5
  fi
  confirm_fix_repo "$script"
  snapshot_pre_fix_repo
  log_dir="$LOG_DIR"
  [[ -z "$log_dir" ]] && log_dir="$DEST/.install-logs"
  case "$log_dir" in /*) ;; *) log_dir="$DEST/$log_dir" ;; esac
  mkdir -p "$log_dir"
  ts="$(date -u +%Y%m%dT%H%M%SZ)"
  log_file="$log_dir/fix-repo-$ts.log"
  echo ""
  step "Running fix-repo: $script"
  step "Log: $log_file"
  {
    echo "# fix-repo log"
    echo "# started:  $(date -u +%Y-%m-%dT%H:%M:%SZ)"
    echo "# script:   $script"
    echo "# dest:     $DEST"
    echo "# os:       $(uname -s 2>/dev/null || echo unknown)"
    echo "# shell:    bash ${BASH_VERSION:-unknown}"
    echo "# uname:    $(uname -a 2>/dev/null || echo unknown)"
    echo "# cwd:      $(pwd)"
    echo "# rollback: on-failure=$ROLLBACK_ON_FIX_FAIL  full=$FULL_ROLLBACK"
    echo "# ──────────────────────────────────────────────────────────"
  } > "$log_file"
  set +e
  case "$script" in
    *.ps1)
      if command -v pwsh >/dev/null 2>&1; then
        pwsh -NoProfile -ExecutionPolicy Bypass -File "$script" 2>&1 | tee -a "$log_file"
        rc=${PIPESTATUS[0]}
      elif command -v powershell >/dev/null 2>&1; then
        powershell -NoProfile -ExecutionPolicy Bypass -File "$script" 2>&1 | tee -a "$log_file"
        rc=${PIPESTATUS[0]}
      else
        set -e
        err "--run-fix-repo: neither pwsh nor powershell found in PATH."
        exit 5
      fi
      ;;
    *)
      bash "$script" 2>&1 | tee -a "$log_file"
      rc=${PIPESTATUS[0]}
      ;;
  esac
  set -e
  echo "# exit: $rc  finished: $(date -u +%Y-%m-%dT%H:%M:%SZ)" >> "$log_file"
  if $SHOW_FIX_REPO_LOG; then
    echo ""
    echo "─── fix-repo log: $log_file ─────────────────────────────"
    cat "$log_file"
    echo "─── end of log ──────────────────────────────────────────"
  fi
  if [[ "$rc" -ne 0 ]]; then
    err "fix-repo failed (exit $rc) — see $log_file"
    $ROLLBACK_ON_FIX_FAIL && perform_rollback "$log_file"
    exit 5
  fi
  ok "fix-repo completed (log: $log_file)"
}
if ! $DRY_RUN && $RUN_FIX_REPO; then run_fix_repo; fi

# ── Summary ───────────────────────────────────────────────────────
echo ""
echo "════════════════════════════════════════════════════════"
[[ $COPIED -gt 0 ]] && ok "$COPIED folder(s) processed"
[[ $WROTE_NEW -gt 0 ]] && ok "$WROTE_NEW new file(s)"
[[ $OVERWROTE -gt 0 ]] && ok "$OVERWROTE file(s) overwritten"
[[ $SKIPPED_FILES -gt 0 ]] && warn "$SKIPPED_FILES file(s) skipped"
[[ $SKIPPED_FOLDERS -gt 0 ]] && warn "$SKIPPED_FOLDERS folder(s) missing in source"
$DRY_RUN && warn "DRY-RUN — no changes written"
echo ""
echo "  Source:      $REPO @ $REF"
echo "  Destination: $DEST"
echo "  Folders:     ${FOLDERS[*]}"
echo ""
echo "════════════════════════════════════════════════════════"
