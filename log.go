package will_log

import (
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"os"
	"time"
)

var (
	defaultTSFormat = log.TimestampFormat(
		func() time.Time { return time.Now().Local() },
		"2006-01-02 15:04:05.000",
	)
)

type Config struct {
	Level  *AllowedLevel
	Format *AllowedFormat
}

type AllowedLevel struct {
	s string
	o level.Option
}

func (l *AllowedLevel) String() string {
	return l.s
}

func (l *AllowedLevel) Set(s string) error {
	switch s {
	case "debug", "d":
		l.o = level.AllowDebug()
	case "info", "i":
		l.o = level.AllowInfo()
	case "warn", "w":
		l.o = level.AllowWarn()
	case "error", "e":
		l.o = level.AllowError()
	default:
		return fmt.Errorf("unrecognized log level %q", s)
	}
	l.s = s
	return nil
}

type AllowedFormat struct {
	s        string
	TSFormat log.Valuer
}

func (a *AllowedFormat) String() string {
	return a.s
}

func (a *AllowedFormat) Set(s string) error {
	switch s {
	case "lf", "json", "logfmt", "j":
		if s == "lf" {
			s = "logfmt"
		}
		if s == "j" {
			s = "json"
		}
		a.s = s
	default:
		return fmt.Errorf("unrecognized log format %q", s)
	}
	return nil
}

func (a *AllowedFormat) SetTSDefault(v log.Valuer) {
	a.TSFormat = v
}

func New(config *Config) log.Logger {
	var l log.Logger
	if config.Format != nil && config.Format.s == "json" {
		l = log.NewJSONLogger(log.NewSyncWriter(os.Stderr))
	} else {
		l = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	}
	if config.Format != nil && config.Format.TSFormat == nil {
		config.Format.TSFormat = defaultTSFormat
	}
	if config.Format == nil {
		config.Format = &AllowedFormat{
			TSFormat: defaultTSFormat,
		}
	}

	if config.Level != nil {
		l = log.With(l, "timestamp", config.Format.TSFormat, "caller", log.Caller(5))
		l = level.NewFilter(l, config.Level.o)
	} else {
		l = log.With(l, "timestamp", config.Format.TSFormat, "caller", log.DefaultCaller)
	}
	return l
}
