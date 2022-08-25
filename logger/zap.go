package logger

import (
	"runtime/debug"
	"time"

	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"go.uber.org/zap"
)

// NewZapLogger will return a new production logger backed by zap
func NewZaplogger() (Logger, error) {
	conf := zap.NewProductionConfig()
	conf.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	conf.DisableCaller = true
	conf.DisableStacktrace = true
	zapLogger, err := conf.Build(zap.AddCaller(), zap.AddCallerSkip(1), zap.WrapCore(zapCore))

	return zpLg{
		lg: zapLogger.Sugar(),
	}, err
}

func zapCore(c zapcore.Core) zapcore.Core {
	// lumberjack.Logger is already safe for concurrent use, so we don't need to
	// lock it.
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./sirchat.log",
		MaxSize:    50, // megabytes
		MaxBackups: 30,
		MaxAge:     28, // days
	})

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.DebugLevel,
	)
	cores := zapcore.NewTee(c, core)

	return cores
}

type zpLg struct {
	lg *zap.SugaredLogger
}

func (l zpLg) Debug(msg string, keyValues ...interface{}) {
	l.lg.With("date", time.Now().UTC().Format(time.RFC3339), "stacktrace", string(debug.Stack())).Debugw(msg, keyValues...)
}

func (l zpLg) Info(msg string, keyValues ...interface{}) {
	l.lg.With("date", time.Now().UTC().Format(time.RFC3339)).Infow(msg, keyValues...)
}

func (l zpLg) Warn(msg string, keyValues ...interface{}) {
	l.lg.With("date", time.Now().UTC().Format(time.RFC3339)).Warnw(msg, keyValues...)
}

func (l zpLg) Error(msg string, keyValues ...interface{}) {
	l.lg.With("date", time.Now().UTC().Format(time.RFC3339), "stacktrace", string(debug.Stack())).Errorw(msg, keyValues...)
}

func (l zpLg) ErrorWithoutSTT(msg string, keyValues ...interface{}) {
	l.lg.With("date", time.Now().UTC().Format(time.RFC3339)).Errorw(msg, keyValues...)
}

func (l zpLg) Fatal(msg string, keyValues ...interface{}) {
	l.lg.With("data", time.Now().UTC().Format(time.RFC3339), "stacktrace", string(debug.Stack())).Fatalw(msg, keyValues...)
}
