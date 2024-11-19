package glog

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type TimestampHook struct{}

func (t TimestampHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	e.Time("datetime", time.Now().UTC())
}

func NewLogger() zerolog.Logger {
	logger := zerolog.New(os.Stdout)
	logger = logger.Hook(TimestampHook{})
	return logger
}
