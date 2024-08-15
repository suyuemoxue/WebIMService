package utils

import (
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
	"time"
)

type LogrusLogger struct {
	Log *logrus.Logger
}

func (l *LogrusLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l *LogrusLogger) Info(ctx context.Context, s string, args ...interface{}) {
	l.Log.WithContext(ctx).Infof(s, args...)
}

func (l *LogrusLogger) Warn(ctx context.Context, s string, args ...interface{}) {
	l.Log.WithContext(ctx).Warnf(s, args...)
}

func (l *LogrusLogger) Error(ctx context.Context, s string, args ...interface{}) {
	l.Log.WithContext(ctx).Errorf(s, args...)
}

func (l *LogrusLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, rows := fc()
	if err != nil {
		l.Log.WithContext(ctx).WithError(err).WithFields(logrus.Fields{
			"rows": rows,
			"sql":  sql,
		}).Error("GORM query error")
	} else {
		l.Log.WithContext(ctx).WithFields(logrus.Fields{
			"rows": rows,
			"sql":  sql,
		}).Debug("GORM query")
	}
}
