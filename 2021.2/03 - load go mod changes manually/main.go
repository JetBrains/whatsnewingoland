package main

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func _() {
	_ = logrus.Level(0)
	_ = zap.AtomicLevel{}
}
