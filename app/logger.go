package app

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"

	"fmt"
	cfg "github.com/duhruh/blog/config"
	"github.com/fatih/color"
)

// This is where we define our application logger
// here we initialize the logger to only output to
// stdout
func NewLogger(c Config) log.Logger {

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = newColorLogger(logger)
	logger = level.NewFilter(logger, c.LogOption())
	logger = log.With(
		logger,
		"timestamp", log.DefaultTimestampUTC,
		"environment", c.Environment(),
		"gitCommit", cfg.GitCommit,
		"version", cfg.Version,
		"buildNumber", cfg.BuildNumber,
		"buildTime", cfg.BuildTime,
		"caller", log.DefaultCaller,
	)

	level.Info(logger).Log("message", "application booting")

	return logger
}

type colorLogger struct {
	next log.Logger
}

func newColorLogger(l log.Logger) log.Logger {
	return colorLogger{next: l}
}

func (l colorLogger) Log(args ...interface{}) error {
	yellow := color.New(color.FgHiYellow).SprintFunc()
	var key []interface{}
	for v, k := range args {
		if (v % 2) == 1 {
			key = append(key, k)
			continue
		}

		key = append(key, fmt.Sprintf("%s", yellow(k)))
	}

	return l.next.Log(key...)
}
