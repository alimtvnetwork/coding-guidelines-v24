# Strictly Avoid

Items in this file MUST NEVER be suggested, recommended, or asked about again.

## readme.txt generator

- ❌ NEVER suggest making the 3-word prefix configurable. It is hard-coded `let's start now`.
- ❌ NEVER suggest a curated/random message list.
- ❌ NEVER suggest "skip if same day", "write only if missing", or any idempotency variant. The script ALWAYS rewrites on every invocation.
- ❌ NEVER suggest local-machine time or any timezone other than `Asia/Kuala_Lumpur`. Timestamp is ALWAYS Malaysia time.
- ❌ NEVER ask the user to re-confirm any of the above.

Locked behavior: prefix `let's start now`, timestamp in `Asia/Kuala_Lumpur` formatted `dd-MMM-yyyy hh:mm:ss tt`, unconditional overwrite of `readme.txt` at repo root.
