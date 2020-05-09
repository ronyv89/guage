package tools

import "os/exec"

// ExecCommand executes a shell command
func ExecCommand(path string, command string, args ...string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = path
	out, err := cmd.Output()
	return out, err
}
