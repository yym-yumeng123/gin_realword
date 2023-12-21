package logger

import (
	"context"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
}

func New(ctx context.Context) *logrus.Entry {
	return logrus.New().WithContext(ctx)
}
