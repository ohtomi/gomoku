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
		port         int
		filename     string
		cors         bool
		verbose      bool
		errorNoMatch bool
	)

	flags := flag.NewFlagSet("run", flag.ContinueOnError)
	flags.Usage = func() {
		c.Ui.Error(c.Help())
	}

	flags.IntVar(&port, "port", 8080, "")
	flags.IntVar(&port, "p", 8080, "")
	flags.StringVar(&filename, "file", "gomoku.yml", "")
	flags.StringVar(&filename, "f", "gomoku.yml", "")
	flags.BoolVar(&cors, "cors", false, "")
	flags.BoolVar(&verbose, "verbose", false, "")
	flags.BoolVar(&errorNoMatch, "error-no-match", false, "")

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

	if err := server.StartHttpServer(fmt.Sprintf(":%d", port), config, cors, verbose, errorNoMatch); err != nil {
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
  --cors      Enable CORS suppport. By default, false.
  --verbose   Print verbosely. By default, false.
  --error-no-match
              Respond 500 internal server error. By default, false (= respond 200 OK).
`
	return strings.TrimSpace(helpText)
}
