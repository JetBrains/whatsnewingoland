package logrus

import (
	"github.com/sirupsen/logrus"
)

type Adapter struct {}

func (a *Adapter) Info(message string, args ...string) error {
	logrus.Infof(message, args)
	return nil
}
