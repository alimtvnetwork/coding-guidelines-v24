package errcmd

import "os/exec"

// buildPowerShellCommand constructs a secure, non-interactive PowerShell invocation.
func buildPowerShellCommand(script string) *exec.Cmd {
	return exec.Command(
		"powershell.exe",
		"-NoProfile",
		"-NonInteractive",
		"-ExecutionPolicy",
		"Bypass",
		"-Command",
		script,
	)
}
