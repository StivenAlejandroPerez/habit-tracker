package logger

import (
	"context"
	"os"

	"github.com/labstack/gommon/random"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

const (
	prefixKey = "Prefix"
	loggerKey = "Logger"
)

func GetLogger(ctx context.Context) *logrus.Entry {
	log, ok := ctx.Value(loggerKey).(*logrus.Entry)
	if !ok {
		return New("-")
	}

	return log
}

func GetPrefix(logger *logrus.Entry) string {
	return logger.Data[prefixKey].(string)
}

func New(prefix string) *logrus.Entry {
	logger := logrus.Logger{
		Out:   os.Stdout,
		Level: logrus.InfoLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02T15:04:05.1483386-00:00",
			LogFormat:       "[%lvl%] [%Prefix%] %msg%\n",
		},
	}

	if prefix == "-" {
		prefix = random.New().String(32, random.Alphanumeric)
	}

	return logger.WithFields(
		logrus.Fields{
			prefixKey: prefix,
		},
	)
}

func NewContext() context.Context {
	return SetLogger(context.Background(), New("-"))
}

func SetLogger(ctx context.Context, logger *logrus.Entry) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}
