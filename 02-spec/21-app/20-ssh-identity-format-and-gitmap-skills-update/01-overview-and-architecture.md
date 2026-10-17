# Architecture Specification: SSH Identity Format, Git Configuration & GitMap LLM Training

> Version: 1.0.0
> Date: 2026-10-17
> Status: Normative
> Scope: Category `27-muse-prompts`, Onboarding Protocol, GitMap Skills & Fleet Synchronization

---

## 1. Executive Summary

During autonomous onboarding (Step 2 of the Muse Onboarding Ritual), if GitHub authentication is not present and Option (b) (SSH Key Generation) is chosen, the generated keypair and Git author identity must adhere to a standardized, deterministic format encoding the operator email, Muse sequence number, and operator name.

### Identity Format Token
```text
<email>-muse<muse-sequence>-<operator-name>
```

### Concrete Examples
- `alex.tone@company.com-muse-10-opRoden`
- `alex.john@company.com-muse-10-opRoden`

---

## 2. Interactive Operator Prompting Protocol

If any of the three identity components are missing or uninitialized in the runtime environment:
1. `<email>`: Target email address
2. `<muse-sequence>`: Numeric sequence identifier of the Muse instance (e.g. `10`, `01`, `02`)
3. `<operator-name>`: Human operator identifier (e.g. `opRoden`, `opAlex`)

The agent **MUST STOP AND ASK** the owner/operator for the missing values before generating the key:
> *"Please provide the required SSH identity variables: (1) your email address, (2) Muse sequence number, and (3) operator name (format: `<email>-muse<sequence>-<operator-name>`)."*

---

## 3. Shell & Git Command Sequence

Once the three variables are confirmed, the agent executes:

```bash
# 1. Configure Git global user identity
git config --global user.email "<email>"
git config --global user.name "<operator-name> (Muse <muse-sequence>)"

# 2. Start ssh-agent
eval "$(ssh-agent -s)"

# 3. Generate Ed25519 keypair with exact comment token
ssh-keygen -t ed25519 -C "<email>-muse<muse-sequence>-<operator-name>" -f ~/.ssh/id_ed25519 -N ""

# 4. Add key to agent
ssh-add ~/.ssh/id_ed25519

# 5. Output public key for operator / GitHub registration
cat ~/.ssh/id_ed25519.pub

# 6. Register key to GitHub via CLI if authorized, or present for web entry
gh ssh-key add ~/.ssh/id_ed25519.pub --title "<email>-muse<muse-sequence>-<operator-name>" || true

# 7. Verify authentication against GitHub
ssh -T git@github.com
```

---

## 4. GitMap Skills & LLM Training Documentation

The GitMap skills (`.agents/skills/gitmap/skill.md` and `.cursor/skills/gitmap/skill.md`) must document:
1. `gitmap llm train` / `gitmap train`: Core curriculum and documentation feeder.
2. `gitmap llm train muse [--days <N>]`: On-demand memory ingest from `repo-cache/12-muse-ai/`, diffing against watermark commit and generating sequential memory reading checklist.
3. Native compiled Go `gitmap agent` subcommands with total ban on `python -c "import sqlite3..."`.
4. Fleet synchronization via `gitmap sync`.
