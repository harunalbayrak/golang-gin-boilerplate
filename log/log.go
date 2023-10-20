package log

import (
	"go.uber.org/zap"
)

type Logger = zap.SugaredLogger // Type alias https://go.dev/ref/spec#Type_declarations

func NewLogger() (logger *Logger) {
	zConfig := zap.NewProductionConfig()
	zConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	z, err := zConfig.Build()
	if err != nil {
		panic("zap logger build error")
	}
	defer z.Sync() // flushes buffer, if any
	logger = z.Sugar()
	return logger
}
