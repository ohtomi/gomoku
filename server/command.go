package server

import (
	"bytes"
	"os/exec"
)

func (c *Command) Execute() (map[string]interface{}, error) {
	stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}
	cmd := c.buildCommand(stdout, stderr)

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	result := map[string]interface{}{
		"Path":       cmd.Path,
		"Args":       cmd.Args,
		"WorkingDir": cmd.Dir,
		"Return":     stdout.String(),
	}

	return result, nil
}

func (c *Command) buildCommand(stdout, stderr *bytes.Buffer) *exec.Cmd {
	var cmd *exec.Cmd

	if len(c.Type) == 0 {
		cmd = exec.Command(c.Path)
	} else {
		cmd = exec.Command(c.Type, c.Path)
	}

	cmd.Stdout = stdout
	cmd.Stderr = stderr

	return cmd
}
