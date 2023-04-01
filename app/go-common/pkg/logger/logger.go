package logger

import (
	"context"

	c "github.com/Calmantara/go-common/pkg/context"
	"go.uber.org/zap"
)

type LogKey string

const (
	corrId LogKey = "X-Correlation-ID"
)

var (
	sugar *zap.SugaredLogger
)

func (c LogKey) String() string {
	return string(c)
}

func init() {
	conf := zap.NewProductionConfig()
	conf.Encoding = "json"
	conf.DisableStacktrace = true

	logger, _ := conf.Build(
		zap.AddCallerSkip(1),
	// zap.AddCaller(),
	)

	defer logger.Sync() // flushes buffer, if any
	sugar = logger.Sugar()
}

func Info(ctx context.Context, template string, args ...interface{}) {
	_, cid := c.GetCorrelationID(ctx)
	args = append(args, corrId.String(), cid)
	if len(args)%2 != 0 {
		sugar.Error("invalid args input")
	}
	sugar.Infow(template, args...)
}

func Error(ctx context.Context, template string, args ...interface{}) {
	_, cid := c.GetCorrelationID(ctx)
	args = append(args, corrId.String(), cid)
	if len(args)%2 != 0 {
		sugar.Error("invalid args input")
	}
	sugar.Errorw(template, args...)
}
