# 04 — Examples

## PowerShell

```powershell
# Toggle current visibility
.\visibility-change.ps1

# Force public (will prompt if going from private)
.\visibility-change.ps1 -Visible pub

# Force public, no prompt (CI use)
.\visibility-change.ps1 -Visible pub -Yes

# Force private
.\visibility-change.ps1 -Visible pri

# Dry-run a toggle
.\visibility-change.ps1 -DryRun

# Help
.\visibility-change.ps1 -Help
```

## Bash

```bash
# Toggle current visibility
./visibility-change.sh

# Force public (will prompt if going from private)
./visibility-change.sh --visible pub

# Force public, no prompt
./visibility-change.sh --visible pub --yes

# Force private
./visibility-change.sh --visible pri

# Dry-run
./visibility-change.sh --dry-run

# Help
./visibility-change.sh --help
```

## Via the runner

```powershell
.\run.ps1 visibility                 # toggle
.\run.ps1 visibility -Visible pub    # force public
.\run.ps1 visibility -DryRun         # preview
.\run.ps1 help                       # see all sub-commands incl. visibility
```

```bash
./run.sh visibility                  # toggle
./run.sh visibility --visible pri    # force private
./run.sh visibility --dry-run        # preview
./run.sh help                        # see all sub-commands incl. visibility
```

## Sample Output

```
$ ./visibility-change.sh
visibility: public → private (github)

$ ./visibility-change.sh --visible pub
⚠  About to make mahin/movie-cli-v2 PUBLIC on github.
   URL: https://github.com/mahin/movie-cli-v2
   Type 'yes' to continue, anything else aborts: yes
visibility: private → public (github)

$ ./visibility-change.sh --visible pub --dry-run
[dry-run] would change private → public (github)
```
