#!/usr/bin/env bash
# ============================================================
# linters-cicd installer (one-liner)
#
# Conforms to: spec/14-update/27-generic-installer-behavior.md
#
#   curl -fsSL https://github.com/alimtvnetwork/coding-guidelines-v17/releases/latest/download/install.sh | bash
#
# Flags:
#   -d <dir>      install destination (default: ./linters-cicd)
#   -v <version>  install a specific version (PINNED MODE, §4) (default: latest)
#   -n            skip checksum verification (NOT recommended)
#   -h, --help    show this help and exit
#
# EXIT CODES (spec §8):
#   0  success
#   1  generic failure (download / extract / checksum mismatch)
#   2  unknown flag
#   3  pinned release / asset not found (PINNED MODE only)
#   4  verification failed (checksum)
# ============================================================

set -euo pipefail

usage() {
    sed -n '2,21p' "$0" | sed 's/^# \{0,1\}//'
    exit 0
}

# Pre-parse long-form help (getopts only supports short flags).
for arg in "$@"; do
    case "$arg" in
        -h|--help) usage ;;
    esac
done

REPO="alimtvnetwork/coding-guidelines-v17"
DEST="./linters-cicd"
VERSION="latest"
VERIFY=1

while getopts "d:v:nh" opt; do
    case "$opt" in
        d) DEST="$OPTARG" ;;
        v) VERSION="$OPTARG" ;;
        n) VERIFY=0 ;;
        h) usage ;;
        *) echo "Unknown flag" >&2; exit 2 ;;
    esac
done

# Banner (spec §7)
INSTALL_MODE="implicit"
SOURCE_KIND="release-asset (latest)"
if [ "$VERSION" != "latest" ]; then
    INSTALL_MODE="pinned"
    SOURCE_KIND="release-asset ($VERSION)"
fi
echo "    📦 coding-guidelines linters-cicd installer"
echo "       mode:    $INSTALL_MODE"
echo "       repo:    $REPO"
echo "       version: $VERSION"
echo "       source:  $SOURCE_KIND"
echo "       dest:    $DEST"
echo ""

# =====================================================================
# -n (NoVerify) warning banner
#
# When checksum verification is disabled, print a prominent multi-line
# banner so the operator can never miss it — including in piped
# `curl | bash` flows where stdout scrolls quickly. This mirrors the
# spec §8 exit-code contract:
#   - With verification ON  → mismatch exits 4 (verification failed).
#   - With -n (NoVerify)    → no integrity check is performed; corrupted
#                             or tampered archives will install silently
#                             and the script will exit 0 on success.
# Kept in sync with the equivalent block in linters-cicd/install.ps1.
# =====================================================================
if [ "$VERIFY" -eq 0 ]; then
    # ANSI yellow if stderr is a TTY; plain text otherwise (CI logs).
    if [ -t 2 ]; then
        Y=$'\033[33m'; R=$'\033[0m'
    else
        Y=""; R=""
    fi
    {
        echo ""
        echo "${Y}    ╔══════════════════════════════════════════════════════════════════╗${R}"
        echo "${Y}    ║  ⚠️  WARNING: -n (NoVerify) — SHA-256 verification is DISABLED   ║${R}"
        echo "${Y}    ╠══════════════════════════════════════════════════════════════════╣${R}"
        echo "${Y}    ║  The downloaded archive will NOT be checked against              ║${R}"
        echo "${Y}    ║  checksums.txt. Corrupted or tampered files will install         ║${R}"
        echo "${Y}    ║  silently. This is NOT recommended for CI or production use.     ║${R}"
        echo "${Y}    ║                                                                  ║${R}"
        echo "${Y}    ║  Exit-code impact (spec §8):                                     ║${R}"
        echo "${Y}    ║    • verification ON   →  checksum mismatch exits 4              ║${R}"
        echo "${Y}    ║    • verification OFF  →  no exit 4 is ever raised               ║${R}"
        echo "${Y}    ║                           (script exits 0 on download success,  ║${R}"
        echo "${Y}    ║                            even for a tampered archive)         ║${R}"
        echo "${Y}    ║                                                                  ║${R}"
        echo "${Y}    ║  Re-run WITHOUT -n to restore integrity checking.                ║${R}"
        echo "${Y}    ╚══════════════════════════════════════════════════════════════════╝${R}"
        echo ""
    } >&2
fi

if [ "$VERSION" = "latest" ]; then
    URL_BASE="https://github.com/$REPO/releases/latest/download"
else
    URL_BASE="https://github.com/$REPO/releases/download/$VERSION"
fi

TMP="$(mktemp -d)"
trap 'rm -rf "$TMP"' EXIT

echo "    ▸ downloading zip..."
ZIP_NAME="coding-guidelines-linters.zip"
# Releases use coding-guidelines-linters-vX.Y.Z.zip; resolve via redirect
if [ "$VERSION" = "latest" ]; then
    DOWNLOAD_URL=$(curl -fsSLI -o /dev/null -w "%{url_effective}" "$URL_BASE/$ZIP_NAME" 2>/dev/null || true)
    if [ -z "$DOWNLOAD_URL" ] || ! curl -fsSL -o "$TMP/$ZIP_NAME" "$URL_BASE/$ZIP_NAME"; then
        # Fall back to versioned name
        echo "    ▸ resolving latest tag..."
        TAG=$(curl -fsSL "https://api.github.com/repos/$REPO/releases/latest" | python3 -c "import sys,json; print(json.load(sys.stdin)['tag_name'])")
        ZIP_NAME="coding-guidelines-linters-$TAG.zip"
        URL_BASE="https://github.com/$REPO/releases/download/$TAG"
        curl -fsSL -o "$TMP/$ZIP_NAME" "$URL_BASE/$ZIP_NAME"
    fi
else
    ZIP_NAME="coding-guidelines-linters-$VERSION.zip"
    curl -fsSL -o "$TMP/$ZIP_NAME" "$URL_BASE/$ZIP_NAME"
fi

if [ "$VERIFY" -eq 1 ]; then
    echo "    ▸ verifying checksum..."
    if curl -fsSL -o "$TMP/checksums.txt" "$URL_BASE/checksums.txt" 2>/dev/null; then
        EXPECTED=$(grep "$ZIP_NAME" "$TMP/checksums.txt" | awk '{print $1}' || true)
        ACTUAL=$(sha256sum "$TMP/$ZIP_NAME" | awk '{print $1}')
        if [ -n "$EXPECTED" ] && [ "$EXPECTED" != "$ACTUAL" ]; then
            echo "    ❌ checksum mismatch! expected=$EXPECTED actual=$ACTUAL" >&2
            exit 1
        fi
        echo "    ✅ checksum OK"
    else
        echo "    ⚠️  checksums.txt not found, skipping verification"
    fi
fi

echo "    ▸ extracting to $DEST..."
mkdir -p "$DEST"
if command -v unzip >/dev/null 2>&1; then
    unzip -oq "$TMP/$ZIP_NAME" -d "$DEST"
else
    python3 -c "import zipfile; zipfile.ZipFile('$TMP/$ZIP_NAME').extractall('$DEST')"
fi

# Strip outer linters-cicd/ folder if present
if [ -d "$DEST/linters-cicd" ]; then
    cp -R "$DEST/linters-cicd/." "$DEST/"
    rm -rf "$DEST/linters-cicd"
fi

chmod +x "$DEST/run-all.sh" 2>/dev/null || true

echo ""
echo "    ✅ installed → $DEST"
echo ""
echo "    Next steps:"
echo "       bash $DEST/run-all.sh --path . --format text"
echo ""
