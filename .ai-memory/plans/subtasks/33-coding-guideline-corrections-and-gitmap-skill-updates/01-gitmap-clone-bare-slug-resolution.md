# Subtask 01: GitMap Clone Bare Slug Resolution

## Objective
Fix `gitmap clone <repo-name>` to resolve bare repository names against the active GitHub profile or user (`resolveActiveOrProfileOwner()`), forming `https://github.com/<owner>/<repo-name>` before giving up.

## Target Files
- `cli/cmdclone/gh_resolver.go`
- `cli/cmdclone/gh_resolver_test.go`

## Steps
1. In `gh_resolver.go`, update `ResolveRepoSlug`:
   - If `!isFullURLOrPath(cleanTarget)` and not found in store/gh-cli and does not contain `/`, check if `isValidSlugPart(cleanTarget)`.
   - Query `owner := resolveActiveOrProfileOwner()`. If non-empty, return `fmt.Sprintf("https://github.com/%s/%s", owner, cleanTarget)`.
2. Add unit test in `gh_resolver_test.go` asserting bare slug resolution.
3. Verify with `go test ./cli/cmdclone`.
