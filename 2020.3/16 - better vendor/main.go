package main

// Step 1. Compare the display in Project Tool Window of
//		   this module's presence/External Library information
//		   versus the "wni20203/extprintf" one

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func main() {
	err := errors.Errorf("Hello %s!", "World")
	logrus.WithError(err).Infof("Hello %s!", "World")

	sugar := zap.NewExample().Sugar()
	defer func() { _ = sugar.Sync() }()
	sugar.Infof("Hello %s!", "World")
}
