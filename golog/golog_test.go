package golog_test

import (
	"os"
	"testing"

	"github.com/idabgsram/golog-warehouse/golog"
)

func Test_StoreLog(t *testing.T) {
	_ = os.Setenv("GOLOG_REDIS_URL", "redis://user:password@localhost:6379/2?protocol=3")
	_ = os.Setenv("GOLOG_CHANNEL", "#log-channel")
	_ = os.Setenv("GOLOG_USERNAME", "LogExampleServices")
	golog.New()
	golog.Slack.Info("Log Stored!")
}

func Test_StoreSilentLog(t *testing.T) {
	_ = os.Setenv("GOLOG_REDIS_URL", "redis://user:password@localhost:6379/2?protocol=3")
	_ = os.Setenv("GOLOG_CHANNEL", "#log-channel")
	_ = os.Setenv("GOLOG_USERNAME", "Log Test")
	golog.New()
	golog.Slack.Info("Log Stored!")
}
