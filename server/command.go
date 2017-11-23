package server

import (
	"bytes"
	"os/exec"
)

func (c *Command) Execute(conversation *Conversation) error {
	stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}
	cmd := c.buildCommand(stdout, stderr)

	if err := cmd.Run(); err != nil {
		return err
	}

	conversation.CommandResult = &CommandResult{
		"Path":       cmd.Path,
		"Args":       cmd.Args,
		"WorkingDir": cmd.Dir,
		"Return":     stdout.String(),
	}

	return nil
}

func (c *Command) buildCommand(stdout, stderr *bytes.Buffer) *exec.Cmd {
	var cmd *exec.Cmd

	if len(c.Args) == 0 {
		cmd = exec.Command(c.Path)
	} else {
		cmd = exec.Command(c.Path, c.Args...)
	}

	cmd.Stdout = stdout
	cmd.Stderr = stderr

	return cmd
}
