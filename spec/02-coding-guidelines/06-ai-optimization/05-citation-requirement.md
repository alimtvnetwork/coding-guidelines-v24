# Citation & Relative Path Requirement for AI Agents

## Mandatory Citation & Relative Path Rule (CODE RED)

Whenever an AI agent generates code, creates plans (`.lovable/plans/pending/`), breaks tasks into subtasks (`.lovable/plans/subtasks/`), writes memory logs (`.lovable/memory/issues/`), explains design decisions, or enforces standards, it **MUST** cite the specific `spec/` or `.lovable/` markdown file and line/section that justifies the action using **STRICTLY RELATIVE PATHS FROM THE GIT REPOSITORY ROOT**.

### 1. Total Ban on Absolute Paths & `file:///` URIs in Repository Files

- **TOTAL BAN:** NEVER write absolute filesystem paths (e.g. `D:\work\project\...`, `C:\Users\...`, `/home/...`) or absolute URI schemes (`file:///d:/work/...`, `file:///C:/...`) inside markdown plans, subtask files, code comments, citations, or committed repository files.
- **PORTABILITY REQUIREMENT:** All paths and markdown links within repository files MUST be relative to the git root so they work seamlessly across Windows, Linux, macOS, and CI/CD pipelines.

### 2. Concrete Examples

#### ❌ INVALID (Absolute Path / File URI):

```markdown
- [SSH Commands](file:///d:/.../.lovable/spec/commands/01-ssh-commands.md) — Why: Defines required behavior.
- [App Error Docs](file:///d:/.../spec/05-coding-guidelines/04-error-handling.md) — Why: Standards for returning results.
- [gitmap/cmd/ssh_login_install_cmd.go](file:///d:/.../gitmap/cmd/ssh_login_install_cmd.go) — Why: Target file.
```

#### ✅ VALID (Strict Relative Git Path):

```markdown
- [SSH Commands](.lovable/spec/commands/01-ssh-commands.md) — Why: Defines required behavior.
- [App Error Docs](spec/05-coding-guidelines/04-error-handling.md) — Why: Standards for returning results.
- [gitmap/cmd/ssh_login_install_cmd.go](gitmap/cmd/ssh_login_install_cmd.go) — Why: Target file.
```
### Why This is Required

- It prevents agents from blending external training data with this repository's strict conventions.
- It provides human reviewers with an immediate paper trail to verify that the agent followed the house style.

### Examples of Valid Citations

- *"Implementing this as an early return to avoid nesting, per `spec/02-coding-guidelines/01-cross-language/01-zero-nesting.md`."*
- *"Returning a structured error with context, per `spec/03-error-manage/02-error-architecture/01-error-handling-reference.md`."*

### Violations

If an agent enforces a rule (e.g., "Variables must be named X") but cannot cite a spec file to back it up, it has failed the anti-hallucination contract. Human reviewers should reject such suggestions.
If an agent outputs absolute file paths (`file:///` or drive letters) into repository markdown files or plans, or enforces a rule without citing a valid relative spec path, it has failed the anti-hallucination contract.
