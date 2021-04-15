package log_test

import (
	"testing"

	"loc.al/module/pkg/log"
	"loc.al/module/pkg/log/logrus"
	"loc.al/module/pkg/log/zap"
)

func TestLoggerAdapters(t *testing.T) {
	tests := []struct {
		name string
		adapter log.Do
		message string
	}{
		{"logrus", &logrus.Adapter{}, "Hello, World!"},
		{"zap", &zap.Adapter{}, "Hello, World!"},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			_ = test.adapter.Info(test.message)
		})
	}
}
