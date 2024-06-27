package logger

import (
	"context"
	"log/slog"
	"os"
	"runtime"
	"time"

	"github.com/urodstvo/messenger-server/libs/config"
)

type Logger interface {
	Info(input string, fields ...any)
	Error(input string, fields ...any)
	Debug(input string, fields ...any)
	Warn(input string, fields ...any)

	WithComponent(name string) Logger
	GetSlog() *slog.Logger
}

type Log struct {
	log     *slog.Logger
	service string
}

type LoggerOptions struct {
	Service   string
	Level     slog.Level
	AddSource bool
	IsJSON    bool
}

func NewFx(opts LoggerOptions) func(config config.Config) Logger {
	return func(config config.Config) Logger {
		return New(
			LoggerOptions{
				Service:   opts.Service,
				AddSource: opts.AddSource,
				Level:     opts.Level,
				IsJSON:    opts.IsJSON,
			},
		)
	}
}

func New(opts LoggerOptions) Logger {
	options := &slog.HandlerOptions{
		AddSource: opts.AddSource,
		Level:     opts.Level,
	}

	var handler slog.Handler

	if opts.IsJSON {
		handler = slog.NewJSONHandler(os.Stdout, options)
	} else {
		handler = slog.NewTextHandler(os.Stdout, options)
	}

	log := slog.New(handler)

	if opts.Service != "" {
		log = log.With(slog.String("service", opts.Service))
	}

	return &Log{
		log:     log,
		service: opts.Service,
	}
}

func (c *Log) handle(level slog.Level, input string, fields ...any) {
	var pcs [1]uintptr
	runtime.Callers(3, pcs[:])
	r := slog.NewRecord(time.Now(), level, input, pcs[0])
	for _, f := range fields {
		r.Add(f)
	}
	_ = c.log.Handler().Handle(context.Background(), r)
}

func (c *Log) Info(input string, fields ...any) {
	c.handle(slog.LevelInfo, input, fields...)
}

func (c *Log) Warn(input string, fields ...any) {
	c.handle(slog.LevelWarn, input, fields...)
}

func (c *Log) Error(input string, fields ...any) {
	c.handle(slog.LevelError, input, fields...)

}

func (c *Log) Debug(input string, fields ...any) {
	c.handle(slog.LevelDebug, input, fields...)
}

func (c *Log) WithComponent(name string) Logger {
	return &Log{
		log:     c.log.With(slog.String("component", name)),
		service: c.service,
	}
}

func (c *Log) GetSlog() *slog.Logger {
	return c.log
}
