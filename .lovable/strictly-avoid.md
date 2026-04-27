# Strictly Avoid

Items in this file MUST NEVER be suggested, recommended, asked about, or built again.

## readme.txt timestamp generator — TOTAL BAN

🔴 **NEVER build, suggest, propose, design, spec, or even mention any feature that writes a timestamp / date / time / "Malaysia-formatted" content into `readme.txt` (or any other file).**

This includes — but is not limited to:
- ❌ A `refresh-readme.ps1` / `refresh-readme.sh` / any script that writes time into readme.txt
- ❌ A `readme` sub-command on `run.ps1` / `run.sh` that touches readme.txt timestamps
- ❌ An npm script (`npm run refresh-readme`, etc.) that writes time into readme.txt
- ❌ Hooking timestamp-writing into `npm run sync` or any other workflow
- ❌ Hard-coded prefix variants (`let's start now`, etc.), configurable prefix, curated lists, random phrases
- ❌ Any timezone discussion (Asia/Kuala_Lumpur, UTC, local) tied to readme.txt
- ❌ Any 12-hr / 24-hr / `dd-MMM-yyyy` format discussion tied to readme.txt
- ❌ Any idempotency variant (always rewrite, skip if same day, write if missing)
- ❌ "Instructions" / "how it works" / "how to run" / "how to test" sections about any such generator
- ❌ Asking clarifying questions about any of the above
- ❌ Offering it as a follow-up, alternative, or "while we're at it" suggestion

**If the user asks for this feature again, do nothing except acknowledge that this entry forbids it. Do not negotiate. Do not propose a "smaller" version. Do not ask "did you mean X". Just stop.**

The only acceptable interaction with `readme.txt` is a one-shot manual edit when the user explicitly types the exact content they want in that turn.

**Why:** User has rejected this feature, the suggestion of this feature, the discussion of this feature, and the documentation of this feature multiple times across sessions, with escalating frustration. Re-raising it is a hard failure.
