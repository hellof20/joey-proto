package logger

import (
	"context"

	"github.com/sirupsen/logrus"
)

type loggerKey struct{}

func NewContext(ctx context.Context, logger *logrus.Entry) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

func FromContext(ctx context.Context) *logrus.Entry {
	if ctx == nil {
		return Log.WithFields(logrus.Fields{})
	}

	if l, ok := ctx.Value(loggerKey{}).(*logrus.Entry); ok {
		return l
	}

	return Log.WithFields(logrus.Fields{})
}

func WithContext(ctx context.Context, fields logrus.Fields) context.Context {
	logger := FromContext(ctx).WithFields(fields)
	return NewContext(ctx, logger)
}
