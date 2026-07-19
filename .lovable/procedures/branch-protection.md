# Promoting checks to REQUIRED on `main`

Turn a passing workflow into a merge blocker on `main`. Backlog item 1 (promote `slides-visual` to required) tracks the two remaining slides checks.

## Preconditions

- Repo admin permission on GitHub.
- `gh` CLI logged in as an admin: `gh auth status`.
- The workflow has been green on `main` for at least one full release cycle. Promoting a flaky check blocks every PR.

## Procedure

1. Regenerate the machine-readable payload from the checked-in expectation file.
   ```bash
   node scripts/print-required-checks.mjs --json
   ```
   Copy the `contexts` array from the output.

2. Add the new check name to `.github/branch-protection.expected.json` under `required`, and remove it from `desired-but-not-yet-required` if present. Commit that change with the same PR that flips branch protection.

3. Apply the branch-protection change. Replace `{owner}/{repo}` and paste the fresh contexts array.
   ```bash
   gh api -X PATCH "repos/{owner}/{repo}/branches/main/protection/required_status_checks" \
     -f strict=true \
     -F 'contexts=["Linter Scripts Validation","Codegen (Rule 9 inverted-fields)","Spec Cross-Reference Validation","Sync Drift Check (all generated files)","Validate version.json schema","visual"]'
   ```
   `strict=true` means PRs must be up-to-date with `main` before merge (blocks stale-branch bypasses).

4. Verify.
   ```bash
   gh api "repos/{owner}/{repo}/branches/main/protection/required_status_checks" \
     --jq '.contexts'
   ```
   Output should exactly match `required` from the expectation file.

5. Re-run the enumerator to confirm no drift.
   ```bash
   node scripts/print-required-checks.mjs --check
   ```

## Rolling back a required check

If a newly-required check turns flaky and starts blocking merges:

```bash
gh api -X PATCH "repos/{owner}/{repo}/branches/main/protection/required_status_checks" \
  -f strict=true \
  -F 'contexts=[ ... same list minus the flaky check ... ]'
```

Then move the check back to `desired-but-not-yet-required` in `.github/branch-protection.expected.json` with a `reason` explaining what broke and what has to be true before re-promotion.

## Why this file exists

The exact `gh api` payload was never captured, so the "promote to required" backlog item kept getting deferred every release. Anyone with admin can now execute this in under a minute.
