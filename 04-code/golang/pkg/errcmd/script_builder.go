package errcmd

import (
	"os"
	"os/exec"
	"runtime"
	"strings"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
)

// ShellType designates the target interpreter shell.
type ShellType string

const (
	ShellPowerShell ShellType = "powershell"
	ShellBash       ShellType = "bash"
	ShellSh         ShellType = "sh"
	ShellCmd        ShellType = "cmd"
)

// ScriptBuilder constructs cross-platform executable shell commands.
type ScriptBuilder struct {
	shell   ShellType
	lines   []string
	env     map[string]string
	workDir string
}

// NewScriptBuilder creates an empty builder defaulting to the host OS shell.
func NewScriptBuilder() *ScriptBuilder {
	return &ScriptBuilder{
		shell: AutoDetectShell(),
		lines: make([]string, 0),
		env:   make(map[string]string),
	}
}

// AutoDetectShell identifies the appropriate native shell for the runtime OS.
func AutoDetectShell() ShellType {
	if runtime.GOOS == "windows" {
		return ShellPowerShell
	}

	return ShellBash
}

// SetShell explicitly assigns the target shell.
func (b *ScriptBuilder) SetShell(shell ShellType) *ScriptBuilder {
	b.shell = shell

	return b
}

// AddLine appends a single command script line.
func (b *ScriptBuilder) AddLine(line string) *ScriptBuilder {
	if line != "" {
		b.lines = append(b.lines, line)
	}

	return b
}

// AddLines appends multiple command script lines.
func (b *ScriptBuilder) AddLines(lines ...string) *ScriptBuilder {
	for _, line := range lines {
		b.AddLine(line)
	}

	return b
}

// SetEnv registers an environment variable.
func (b *ScriptBuilder) SetEnv(key, val string) *ScriptBuilder {
	b.env[key] = val

	return b
}

// SetWorkDir sets the execution working directory.
func (b *ScriptBuilder) SetWorkDir(dir string) *ScriptBuilder {
	b.workDir = dir

	return b
}

// ScriptText joins all script lines using appropriate shell newlines.
func (b *ScriptBuilder) ScriptText() string {
	sep := "\n"
	if b.shell == ShellPowerShell || b.shell == ShellCmd {
		sep = "; "
	}

	return strings.Join(b.lines, sep)
}

// BuildCommand compiles the configured script into an *exec.Cmd.
func (b *ScriptBuilder) BuildCommand() (*exec.Cmd, *appfault.AppError) {
	if len(b.lines) == 0 {
		return nil, appfault.New(errtype.Validation, "cannot build command with empty script lines")
	}

	var cmd *exec.Cmd
	switch b.shell {
	case ShellPowerShell:
		cmd = buildPowerShellCommand(b.ScriptText())
	case ShellCmd:
		cmd = exec.Command("cmd.exe", "/c", b.ScriptText())
	default:
		cmd = buildBashCommand(b.shell, b.ScriptText())
	}

	b.applyEnvAndDir(cmd)

	return cmd, nil
}

// applyEnvAndDir sets working directory and environment on the exec.Cmd.
func (b *ScriptBuilder) applyEnvAndDir(cmd *exec.Cmd) {
	if b.workDir != "" {
		cmd.Dir = b.workDir
	}

	if len(b.env) > 0 {
		cmd.Env = os.Environ()
		for k, v := range b.env {
			cmd.Env = append(cmd.Env, k+"="+v)
		}
	}
}
