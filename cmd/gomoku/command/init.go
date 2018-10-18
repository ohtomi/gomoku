package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/ohtomi/gomoku/server"
)

type InitCommand struct {
	Meta
}

func (c *InitCommand) Run(args []string) int {
	var (
		dirname string
	)

	flags := flag.NewFlagSet("init", flag.ContinueOnError)
	flags.Usage = func() {
		c.Ui.Error(c.Help())
	}

	if err := flags.Parse(args); err != nil {
		return 1
	}

	if len(flags.Args()) < 1 {
		c.Ui.Error("DIR not specified")
		return 1
	}

	dirname = flags.Args()[0]
	if err := server.CreateScaffold(dirname); err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	return 0
}

func (c *InitCommand) Synopsis() string {
	return fmt.Sprintf("Create gomoku project under the specified directory")
}

func (c *InitCommand) Help() string {
	helpText := `usage: gomoku init DIR
`
	return strings.TrimSpace(helpText)
}
