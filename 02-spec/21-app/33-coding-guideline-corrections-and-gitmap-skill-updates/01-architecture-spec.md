# Architecture Specification: Coding Guideline Corrections & GitMap Skill Updates

## 1. Executive Summary

This architecture specification details the corrective actions across `coding-guidelines` and `gitmap`:
1. **GitMap Clone Bare Slug Resolution:** When a user executes `gitmap clone <repo-name>` without specifying a full URL or owner prefix, GitMap resolves the repository against the active GitHub profile or user (`resolveActiveOrProfileOwner()`), forming `https://github.com/<owner>/<repo-name>` before falling back to local search.
2. **Context Menu & Show Log Error Resilience:**
   - In folder background context menus, add the Installer Manager UI directly.
   - When `gitmap show-log` is executed in a directory that is not a Git repository (e.g. `D:\work`), instead of terminating with a fatal Go runtime stack trace (`[E9000:EXECUTION] not a git repository: D:\work`), GitMap gracefully detects child repositories, offers an interactive selector, or routes to `gitmap ui` (the repositories explorer).
3. **Show Log UI Modernization:**
   - Add PNG icon embedding and Web App Manifest metadata so Chrome/Edge `--app` mode windows display the authentic GitMap logo rather than the generic browser globe icon.
   - Compact window control buttons (Fullscreen, Refresh, Branch comparison, Close) into sleek icon buttons with hover tooltips.
   - Provide Beyond Compare / VS Code diff invocation options configurable from the UI.
4. **GitMap Skills & In-Engine Help Alignment:**
   - Eliminate uppercase commit module prefixes (`SKILLS`, `CLONER`, `DOCS`, `RELEASE`) in favor of Title Case or lowercase.
   - Remove deprecated cluster commands (`cluster`, `sc`, `ssh`), OS dev clean, and VS Code repair examples.
   - Sync corrections to `d:\work\gitmap\cli\cmd\llm\llm_skill.go`.
5. **Follow Coding Guidelines & Skills Consolidation:**
   - Author detailed `/learn` integration in `follow-coding-guidelines/skill.md`.
   - Audit the remaining skills catalog and establish a structured reduction matrix for user alignment.

---

## 2. Component Design & Behavioral Contracts

### 2.1. GitMap Clone Bare Slug Fallback (`cli/cmdclone/gh_resolver.go`)
- In `ResolveRepoSlug(target string) string`:
  1. If `isFullURLOrPath(target)`, return unchanged.
  2. If cached in SQLite store via `resolveFromStore`, return cached URL.
  3. If resolved via `resolveFromGitHubCLI`, return GitHub URL.
  4. If `isGitHubOwnerRepo(cleanTarget)` (contains owner/repo), return `"https://github.com/" + cleanTarget`.
  5. **New Bare Slug Fallback:** If `isValidSlugPart(cleanTarget)` and does not contain `/`:
     ```go
     owner := resolveActiveOrProfileOwner()
     if len(owner) > 0 && isValidSlugPart(owner) {
         return fmt.Sprintf("https://github.com/%s/%s", owner, cleanTarget)
     }
     ```
  6. Return `cleanTarget`.

### 2.2. Show Log Non-Git Directory Resilience (`cli/cmdlogviewer/server.go`)
- When `findRepoRoot(dir)` fails with `exit status 128` (not a Git repository):
  - Check if child directories contain Git repositories (e.g., scan 1 level deep).
  - If child repositories exist, launch `gitmap ui` (or present child repos list) rather than dumping an `AppError` stack trace to standard error.
  - Print a clean, helpful guide message instead of an execution fault stack trace.

### 2.3. Show Log UI Icon & Header Optimization (`cli/cmdlogviewer/`)
- In `favicon.go` and `ui_assets_html.go`:
  - Provide a 32x32 and 192x192 PNG data URI or endpoint `/icon.png` generated from the GitMap brand SVG.
  - Add `<link rel="icon" type="image/png" sizes="32x32" href="/favicon.png">` and `<link rel="manifest" href="/manifest.json">`.
  - Pack header buttons into a unified icon toolbar group with tooltips.

---

## 3. Adherence to Rules & Invariants
- **AGENTS.md §4:** Strict lowercase file naming.
- **AGENTS.md §5:** Strictly relative Git paths across all docs and specs (`02-spec/...`, `.ai-memory/...`).
- **AGENTS.md §10:** GitMap hyphen-separated commit syntax without colons in message arguments.
- **AGENTS.md §13:** Mandatory Completion Invariant ('No Push = Not Done').
