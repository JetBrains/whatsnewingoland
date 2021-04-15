package zap

import (
	"go.uber.org/zap"
)

type Adapter struct {}

var sugar = zap.NewExample().Sugar()

func (a *Adapter) Info(message string, args ...string) error {
	sugar.Infof(message, args)
	return nil
}
