package server

import (
	"fmt"
	"github.com/mitchellh/cli"
)

type Reporter interface {
	IsEnabled() bool

	Info(string)
	Infof(format string, a ...interface{})
	Warn(string)
	Warnf(format string, a ...interface{})
	Error(string)
	Errorf(format string, a ...interface{})
}

type NoopReporter struct{}

func (r NoopReporter) IsEnabled() bool                        { return false }
func (r NoopReporter) Info(s string)                          {}
func (r NoopReporter) Infof(format string, a ...interface{})  {}
func (r NoopReporter) Warn(s string)                          {}
func (r NoopReporter) Warnf(format string, a ...interface{})  {}
func (r NoopReporter) Error(s string)                         {}
func (r NoopReporter) Errorf(format string, a ...interface{}) {}

type CliReporter struct {
	cli.Ui
}

func (r CliReporter) IsEnabled() bool                        { return true }
func (r CliReporter) Info(s string)                          { r.Ui.Info(s) }
func (r CliReporter) Infof(format string, a ...interface{})  { r.Ui.Info(fmt.Sprintf(format, a...)) }
func (r CliReporter) Warn(s string)                          { r.Ui.Warn(s) }
func (r CliReporter) Warnf(format string, a ...interface{})  { r.Ui.Warn(fmt.Sprintf(format, a...)) }
func (r CliReporter) Error(s string)                         { r.Ui.Error(s) }
func (r CliReporter) Errorf(format string, a ...interface{}) { r.Ui.Error(fmt.Sprintf(format, a...)) }

func NewReporter(ui cli.Ui, enabled bool) Reporter {
	if enabled {
		return CliReporter{ui}
	} else {
		return NoopReporter{}
	}
}
