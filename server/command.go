package server

import (
	"bytes"
	"os/exec"
	"strings"
)

func (c *Command) Execute(conversation *Conversation) error {
	if len(c.Path) == 0 {
		conversation.CommandResult = &CommandResult{}
		return nil
	}

	stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}
	cmd, err := c.buildCommand(stdout, stderr, conversation)
	if err != nil {
		return err
	}

	if err := cmd.Run(); err != nil {
		return err
	}

	conversation.CommandResult = &CommandResult{
		Env:    cmd.Env,
		Path:   cmd.Path,
		Args:   cmd.Args,
		Dir:    cmd.Dir,
		Stdout: strings.TrimRight(stdout.String(), "\n"),
		Stderr: strings.TrimRight(stderr.String(), "\n"),
	}

	return nil
}

func (c *Command) buildCommand(stdout, stderr *bytes.Buffer, conversation *Conversation) (*exec.Cmd, error) {
	var cmd *exec.Cmd

	if len(c.Args) == 0 {
		cmd = exec.Command(c.Path)
	} else {
		args := make([]string, len(c.Args))
		for i, v := range c.Args {
			applied, err := ApplyTemplateText("args", v, conversation)
			if err != nil {
				return nil, err
			}
			args[i] = applied
		}
		cmd = exec.Command(c.Path, args...)
	}

	if len(c.Env) > 0 {
		env := make([]string, len(c.Env))
		for i, v := range c.Env {
			applied, err := ApplyTemplateText("env", v, conversation)
			if err != nil {
				return nil, err
			}
			env[i] = applied
		}
		cmd.Env = env
	}

	cmd.Stdout = stdout
	cmd.Stderr = stderr

	return cmd, nil
}
