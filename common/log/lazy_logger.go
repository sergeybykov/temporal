package log

import (
	"sync"

	"go.temporal.io/server/common/errorcode"
	"go.temporal.io/server/common/log/tag"
)

var _ Logger = (*lazyLogger)(nil)

type (
	lazyLogger struct {
		logger Logger
		tagFn  func() []tag.Tag

		once sync.Once
	}
)

func NewLazyLogger(logger Logger, tagFn func() []tag.Tag) *lazyLogger {
	return &lazyLogger{
		logger: logger,
		tagFn:  tagFn,
	}
}

func (l *lazyLogger) Debug(msg string, tags ...tag.Tag) {
	l.once.Do(l.tagLogger)
	l.logger.Debug(msg, tags...)
}

func (l *lazyLogger) Info(msg string, tags ...tag.Tag) {
	l.once.Do(l.tagLogger)
	l.logger.Info(msg, tags...)
}

func (l *lazyLogger) Warn(msg string, tags ...tag.Tag) {
	l.once.Do(l.tagLogger)
	l.logger.Warn(msg, tags...)
}

func (l *lazyLogger) Error(msg string, tags ...tag.Tag) {
	l.once.Do(l.tagLogger)
	// Ensure error code is present for all error logs
	if !hasErrorCode(tags) {
		tags = append(tags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))
	}
	l.logger.Error(msg, tags...)
}

func (l *lazyLogger) DPanic(msg string, tags ...tag.Tag) {
	l.once.Do(l.tagLogger)
	l.logger.DPanic(msg, tags...)
}

func (l *lazyLogger) Panic(msg string, tags ...tag.Tag) {
	l.once.Do(l.tagLogger)
	l.logger.Panic(msg, tags...)
}

func (l *lazyLogger) Fatal(msg string, tags ...tag.Tag) {
	l.once.Do(l.tagLogger)
	l.logger.Fatal(msg, tags...)
}

func (l *lazyLogger) tagLogger() {
	l.logger = With(l.logger, l.tagFn()...)
}

func hasErrorCode(tags []tag.Tag) bool {
	for _, t := range tags {
		if t.Key() == "error-code" {
			return true
		}
	}
	return false
}
