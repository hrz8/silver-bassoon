package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/hrz8/silver-bassoon/pkg/helper"
)

type LogLevel string

var (
	output = zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}
	logger = zerolog.New(output).With().Timestamp().Logger()
)

const (
	DEBUG LogLevel = "debug"
	INFO  LogLevel = "info"
	WARN  LogLevel = "warn"
	ERROR LogLevel = "error"
)

// Log is a log wrapper at given level on stdout. Usage example:
//
//	logger.Log(logger.INFO, "some message here")
func Log(level LogLevel, msg string, payload ...any) {
	switch level {
	case INFO:
		Info(msg, payload...)
	case WARN:
		Warn(msg, payload...)
	case DEBUG:
		Debug(msg, payload...)
	case ERROR:
		Errorp(msg, payload...)
	default:
		Debug(msg, payload...)
	}
}

// Debug is a log wrapper at debug level on stdout. Usage example:
//
//	logger.Debug("some debug here")
//	logger.Debug("some debug with no payload", nil)
func Debug(msg string, payload ...any) {
	if len(payload) > 0 {
		str := helper.JSONMarshal(payload[0])
		logger.Debug().RawJSON("payload", []byte(str)).Msg(msg)
	} else {
		logger.Debug().Msg(msg)
	}
}

// Info is a log wrapper at info level on stdout. Usage example:
//
//	logger.Info("some info log")
//	logger.Info("some info", map[string]any{"foo": "bar"})
//	logger.Info("some info", &User{Name: "John", Address: &Address{City: "Canberra"}})
func Info(msg string, payload ...any) {
	if len(payload) > 0 {
		str := helper.JSONMarshal(payload[0])
		logger.Info().RawJSON("payload", []byte(str)).Msg(msg)
	} else {
		logger.Info().Msg(msg)
	}
}

// Warn is a log wrapper at warning level on stdout. Usage example:
//
//	logger.Warn("some warning")
func Warn(msg string, payload ...any) {
	if len(payload) > 0 {
		str := helper.JSONMarshal(payload[0])
		logger.Warn().RawJSON("payload", []byte(str)).Msg(msg)
	} else {
		logger.Warn().Msg(msg)
	}
}

// Errorp is a log wrapper at error level on stdout with additional data. Usage example:
//
//	logger.Errorp("some error with payload", map[string]any{"foo": "bar"})
func Errorp(msg string, payload ...any) {
	if len(payload) > 0 {
		str := helper.JSONMarshal(payload[0])
		logger.Error().RawJSON("payload", []byte(str)).Msg(msg)
	} else {
		logger.Error().Msg(msg)
	}
}

// Errorp is a log wrapper at error level on stdout with stack trace output. Usage example:
//
//	logger.Error("some error", errors.New("validation error"))
func Error(msg string, err error) {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	errorLogger := zerolog.New(output).With().Timestamp().Logger()

	errorLogger.Error().Stack().Err(err).Msg(msg)
}

// Fatal is a log wrapper at fatal level on stdout with stack trace output. Usage example:
//
//	logger.Fatal("some error", errors.New("fatal error"))
func Fatal(msg string, err error) {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	errorLogger := zerolog.New(output).With().Timestamp().Logger()

	errorLogger.Fatal().Stack().Err(err).Msg(msg)
}

// JSON is a log wrapper at given level on stdout as raw json formation.
//
// Structured JSON:
//
//	logger.JSON(logger.INFO).Str("foo", "bar").Msg("some log message")
//
// With 'payload' field:
//
//	logger.JSON(
//		logger.DEBUG,
//		map[string]any{"data": "string"}
//	)
//	.Str("foo", "bar")
//	.Msg("some log message")
//
// With struct as 'payload' field:
//
//	logger.JSON(logger.WARN, &User{Name: "John"}).Str("foo", "bar").Msg("some log message")
func JSON(level LogLevel, payload ...any) *zerolog.Event {
	const payloadKey = "payload"

	var str string

	if len(payload) > 0 {
		str = helper.JSONMarshal(payload[0])
	}

	switch level {
	case INFO:
		if str != "" {
			return log.Info().RawJSON(payloadKey, []byte(str))
		}

		return log.Info()
	case WARN:
		if str != "" {
			return log.Warn().RawJSON(payloadKey, []byte(str))
		}

		return log.Warn()
	case ERROR:
		if str != "" {
			return log.Error().RawJSON(payloadKey, []byte(str))
		}

		return log.Error()
	case DEBUG:
		if str != "" {
			return log.Debug().RawJSON(payloadKey, []byte(str))
		}

		return log.Debug()
	default:
		if str != "" {
			return log.Debug().RawJSON(payloadKey, []byte(str))
		}

		return log.Debug()
	}
}
