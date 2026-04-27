#!/usr/bin/env bash
# Writes readme.txt with a fixed greeting + current Malaysia (Asia/Kuala_Lumpur)
# date in dd-MMM-YYYY and time in 12-hour clock format.
set -euo pipefail

OUT="${1:-readme.txt}"
TZ="Asia/Kuala_Lumpur" date +"let's start now %d-%b-%Y %I:%M:%S %p" > "$OUT"
echo "wrote: $OUT"
cat "$OUT"
