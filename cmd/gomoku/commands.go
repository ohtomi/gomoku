package main

import (
	"github.com/mitchellh/cli"
	"github.com/ohtomi/gomoku/cmd/gomoku/command"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:     *meta,
				Version:  Version,
				Revision: GitCommit,
				Name:     Name,
			}, nil
		},
	}
}
