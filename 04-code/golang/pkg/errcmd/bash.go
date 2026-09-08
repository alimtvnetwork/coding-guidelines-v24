package errcmd

import "os/exec"

// buildBashCommand constructs a POSIX compliant bash or sh invocation.
func buildBashCommand(shell ShellType, script string) *exec.Cmd {
	binary := "bash"
	if shell == ShellSh {
		binary = "sh"
	}

	return exec.Command(binary, "-c", script)
}
