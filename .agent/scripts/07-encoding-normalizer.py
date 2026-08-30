#!/usr/bin/env python3
"""
Encoding & Control Character Normalizer.
Scans repository text files, removes corrupt control characters (NUL, BEL, BS, VT, FF, ESC),
and ensures clean UTF-8 text formatting without Byte Order Marks (BOM).

Usage:
  python .lovable/ai-fix-scripts/07-encoding-normalizer.py
"""

import os
import sys

EXTS = ('.md', '.json', '.txt', '.py', '.ps1', '.sh', '.mjs', '.ts', '.tsx', '.go', '.yml', '.yaml', '.css', '.html')
EXCLUDE_DIRS = {'.git', 'node_modules', 'dist', 'build', '.ci-out', 'tmp', '.agent', 'release-artifacts'}

def normalize_file(filepath):
    try:
        with open(filepath, 'rb') as f:
            data = f.read()
    except Exception:
        return False

    original = data
    # Strip UTF-8 BOM if present
    if data.startswith(b'\xef\xbb\xbf'):
        data = data[3:]

    # Remove non-printable control characters (< 32, except tab \t=9, LF \n=10, CR \r=13)
    bad_bytes = [b for b in data if b < 32 and b not in (9, 10, 13)]
    if bad_bytes:
        cleaned = bytearray()
        for b in data:
            if b == 0:
                cleaned.extend(b'0')
            elif b == 7:
                cleaned.extend(b'a')
            elif b == 8:
                cleaned.extend(b'b')
            elif b == 11:
                cleaned.extend(b'v')
            elif b == 12:
                cleaned.extend(b'f')
            elif b == 27:
                cleaned.extend(b'e')
            elif b < 32 and b not in (9, 10, 13):
                pass
            else:
                cleaned.append(b)
        data = bytes(cleaned)

    if data != original:
        with open(filepath, 'wb') as f:
            f.write(data)
        return True
    return False

def main():
    print("Running 07-encoding-normalizer.py...")
    fixed_count = 0
    for root, dirs, files in os.walk('.'):
        dirs[:] = [d for d in dirs if d not in EXCLUDE_DIRS]
        for f in files:
            if f.endswith(EXTS):
                path = os.path.join(root, f)
                if normalize_file(path):
                    print(f"Normalized: {path}")
                    fixed_count += 1

    print(f"Normalized encoding and cleaned control characters in {fixed_count} files.")
    sys.exit(0)

if __name__ == "__main__":
    main()
