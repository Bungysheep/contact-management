package logger

import (
	"os"
	"time"

	"github.com/bungysheep/contact-management/pkg/common/constant/datetimeformat"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// Log zap logger
	Log *zap.Logger
)

// InitLog initializes logger
func InitLog() error {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	logConfig := zap.NewProductionEncoderConfig()
	logConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(datetimeformat.LongDate))
	}

	consoleEncoder := zapcore.NewJSONEncoder(logConfig)

	consoleDebugging := zapcore.AddSync(os.Stdout)
	consoleError := zapcore.AddSync(os.Stderr)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
		zapcore.NewCore(consoleEncoder, consoleError, highPriority),
	)

	Log = zap.New(core)

	return nil
}
