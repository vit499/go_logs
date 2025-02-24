package logger

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	*zerolog.Logger
}

var (
	logger Logger
	once   sync.Once
)

func Get(file string) *Logger {

	once.Do(func() {
		lumberjackLogger := &lumberjack.Logger{
			Filename:   file,
			MaxSize:    10, // megabytes
			MaxBackups: 3,
		}
		consoleWriter := zerolog.New(os.Stdout)
		multi := zerolog.MultiLevelWriter(consoleWriter, lumberjackLogger)
		zeroLogger := zerolog.New(multi).With().Timestamp().Logger()

		// zeroLogger := zerolog.New(os.Stderr).With().Timestamp().Logger()
		loglevel := "info"
		// Set proper loglevel based on config
		switch loglevel {
		case "debug":
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		case "info":
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		case "warn", "warning":
			zerolog.SetGlobalLevel(zerolog.WarnLevel)
		case "err", "error":
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		case "fatal":
			zerolog.SetGlobalLevel(zerolog.FatalLevel)
		case "panic":
			zerolog.SetGlobalLevel(zerolog.PanicLevel)
		default:
			zerolog.SetGlobalLevel(zerolog.InfoLevel) // log info and above by default
		}
		logger = Logger{&zeroLogger}
	})
	return &logger
}
