package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/ohtomi/gomoku/server"
)

type RunCommand struct {
	Meta
}

func (c *RunCommand) Run(args []string) int {
	var (
		port int
	)

	flags := flag.NewFlagSet("run", flag.ContinueOnError)
	flags.Usage = func() {
		c.Ui.Error(c.Help())
	}

	flags.IntVar(&port, "port", 8080, "")
	flags.IntVar(&port, "p", 8080, "")

	if err := flags.Parse(args); err != nil {
		return 1
	}

	if err := server.StartHttpServer(fmt.Sprintf(":%d", port)); err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	return 0
}

func (c *RunCommand) Synopsis() string {
	return fmt.Sprintf("Run HTTP server handling requests according to gomoku.yml")
}

func (c *RunCommand) Help() string {
	helpText := `usage: gomoku run [options...]
Options:
  --port, -p  Port number listened by gomoku HTTP server. By default, 8080.
`
	return strings.TrimSpace(helpText)
}
