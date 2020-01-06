package jbook

import "github.com/sirupsen/logrus"

var logger *logrus.Logger

func init() {
	logger = logrus.New()
}

// SetLogger sets the package-wide logger.
func SetLogger(l *logrus.Logger) {
	logger = l
}
