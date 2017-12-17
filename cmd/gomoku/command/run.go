package command

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ohtomi/gomoku/server"
)

type RunCommand struct {
	Meta
}

func (c *RunCommand) Run(args []string) int {
	var (
		port     int
		filename string
		verbose  bool
	)

	flags := flag.NewFlagSet("run", flag.ContinueOnError)
	flags.Usage = func() {
		c.Ui.Error(c.Help())
	}

	flags.IntVar(&port, "port", 8080, "")
	flags.IntVar(&port, "p", 8080, "")
	flags.StringVar(&filename, "file", "gomoku.yml", "")
	flags.StringVar(&filename, "f", "gomoku.yml", "")
	flags.BoolVar(&verbose, "verbose", false, "")

	if err := flags.Parse(args); err != nil {
		return 1
	}

	config, err := server.NewConfig(filename)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	if err := os.Chdir(filepath.Dir(filename)); err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	if err := server.StartHttpServer(fmt.Sprintf(":%d", port), config, verbose); err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	return 0
}

func (c *RunCommand) Synopsis() string {
	return fmt.Sprintf("Run HTTP server handling requests according to gomoku.yml")
}

func (c *RunCommand) Help() string {
	helpText := `usage: gomoku run [OPTIONS...]
Options:
  --port, -p  Port number listened by gomoku HTTP server. By default, 8080.
  --file, -f  Path to config file. By default, "./gomoku.yml".
  --webui     URL to web UI. e.g. "/console".
  --verbose   Print verbosely. By default, false.
`
	return strings.TrimSpace(helpText)
}
