package will_log

import (
	"github.com/go-kit/log/level"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	logger := New(&Config{})
	if err := logger.Log("hello", "world"); err != nil {
		t.Fatal(err)
	}
}

func TestDefaultConfigWithLevel(t *testing.T) {
	logger := New(&Config{})
	level.Info(logger).Log("info", "with level")
	level.Error(logger).Log("err", "with level")
}

func TestLogFormat(t *testing.T) {
	formatJson := &AllowedFormat{}
	formatJson.Set("j")
	loggerJson := New(&Config{
		Format: formatJson,
	})
	level.Info(loggerJson).Log("format", "json")

	formatLF := &AllowedFormat{}
	formatLF.Set("lf")
	logger := New(&Config{
		Format: formatLF,
	})
	level.Info(logger).Log("format", "log format")
}

func TestLogLevel(t *testing.T) {
	levelErr := &AllowedLevel{}
	levelErr.Set("i")
	logger := New(&Config{
		Level: levelErr,
	})
	level.Debug(logger).Log("ll", "debug")
	level.Info(logger).Log("ll", "info")
	level.Warn(logger).Log("ll", "warn")
	level.Error(logger).Log("ll", "err")
}
