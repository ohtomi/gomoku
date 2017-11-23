package server

import (
	"bytes"
	"os/exec"
	"text/template"
)

func (c *Command) Execute(conversation *Conversation) error {
	stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}
	cmd, err := c.buildCommand(stdout, stderr, conversation)
	if err != nil {
		return err
	}

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

func (c *Command) buildCommand(stdout, stderr *bytes.Buffer, conversation *Conversation) (*exec.Cmd, error) {
	var cmd *exec.Cmd

	if len(c.Args) == 0 {
		cmd = exec.Command(c.Path)
	} else {
		args := make([]string, len(c.Args))
		for i, v := range c.Args {
			buf := &bytes.Buffer{}
			t, err := template.New("args").Parse(v)
			if err != nil {
				return nil, err
			}
			if err := t.Execute(buf, conversation); err != nil {
				return nil, err
			}
			args[i] = buf.String()
		}
		cmd = exec.Command(c.Path, args...)
	}

	cmd.Stdout = stdout
	cmd.Stderr = stderr

	return cmd, nil
}
