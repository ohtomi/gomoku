package command

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

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
		tls          bool
		cert         string
		key          string
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
	flags.BoolVar(&tls, "tls", false, "")
	flags.BoolVar(&tls, "ssl", false, "")
	flags.StringVar(&cert, "cert", "", "")
	flags.StringVar(&key, "key", "", "")
	flags.BoolVar(&errorNoMatch, "error-no-match", false, "")

	if err := flags.Parse(args); err != nil {
		return 1
	}

	if tls && (len(cert) == 0 || len(key) == 0) {
		c.Ui.Warn("certificate file and private key must be given if TLS mode enabled")
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

	go func() {
		ticker := time.Tick(10 * time.Second)
		for _ = range ticker {
			newConfig, err := server.NewConfig(filepath.Base(filename))
			if err != nil {
				c.Ui.Error(err.Error())
				continue
			}
			if config.EqualTo(newConfig) {
				continue
			}
			c.Ui.Output(fmt.Sprintf("Updated configuration [%s]", filename))
			copy(*config, *newConfig)
		}
	}()

	reporter := server.NewReporter(c.Ui)
	if err := server.StartHttpServer(fmt.Sprintf(":%d", port), config, cors, tls, cert, key, errorNoMatch, reporter); err != nil {
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
  --tls       Enable TLS mode. By default, false.
  --cert      Path to certificate file, which must be given if TLS mode enabled.
  --key       Path to private key, which must be given if TLS mode enabled.
  --error-no-match
              Respond 500 internal server error. By default, false (= respond 200 OK).
`
	return strings.TrimSpace(helpText)
}
