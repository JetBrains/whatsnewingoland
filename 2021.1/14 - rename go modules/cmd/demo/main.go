package main

import (
	"renameab.le/module/pkg/log"
	"renameab.le/module/pkg/log/logrus"
	"renameab.le/module/pkg/log/zap"
)

func main() {
	message := "Hello, World!"
	process(&logrus.Adapter{}, message)
	process(&zap.Adapter{}, message)
}

func process(a log.Do, message string) {
	_ = a.Info(message)
}
