# fix-repo fixture

Synthetic input for `fix-repo.sh` / `fix-repo.ps1` verification.

Tokens that MUST be rewritten by default mode (last 2) when current = v18:
- coding-guidelines-v24
- coding-guidelines-v24

Tokens that MUST stay unchanged in default mode:
- coding-guidelines-v24   (only touched by --all)
- coding-guidelines-v24  (only touched by --3 or wider)
- coding-guidelines-v170 (numeric-overflow guard)

URL forms (token must be rewritten, host/owner preserved):
- https://github.com/alimtvnetwork/coding-guidelines-v24/install.sh
- git@github.com:alimtvnetwork/coding-guidelines-v24.git
- ssh://git@github.com/alimtvnetwork/coding-guidelines-v24.git
