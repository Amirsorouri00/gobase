package log

import (
	"errors"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	service "portfolio/services"
)

var logger = log.Logger

// Field adds a field to the logger.
type Field struct {
	Key   string
	Value interface{}
}

// F is a helper function to create a field.
func F(key string, value interface{}) Field {
	return Field{Key: key, Value: value}
}

// Init initializes the logger with the given writer.
func Init(w io.Writer, debug bool) {
	level := zerolog.InfoLevel
	if debug {
		level = zerolog.DebugLevel
	}
	logger = zerolog.New(zerolog.MultiLevelWriter(
		zerolog.ConsoleWriter{Out: os.Stderr},
		w,
	)).With().Timestamp().Logger().Level(level)
}

// Debug logs an err in debug level.
func Debug(err error) {
	logErr(logger.Debug(), err)
}

// Debugf logs a formatted message in debug level.
func Debugf(format string, args ...interface{}) {
	logMsg(logger.Debug(), format, args...)
}

// Info logs an err in info level.
func Info(err error) {
	logErr(logger.Info(), err)
}

// Infof logs a formatted message in info level.
func Infof(format string, args ...interface{}) {
	logMsg(logger.Info(), format, args...)
}

// Warn logs an err in warning level.
func Warn(err error) {
	logErr(logger.Warn(), err)
}

// Warnf logs a formatted message in warning level.
func Warnf(format string, args ...interface{}) {
	logMsg(logger.Warn(), format, args...)
}

// Error logs an err in error level.
func Error(err error) {
	logErr(logger.Error(), err)
}

// Errorf logs a formatted message in error level.
func Errorf(format string, args ...interface{}) {
	logMsg(logger.Error(), format, args...)
}

// Fatal logs an err in fatal level and terminates the program.
func Fatal(err error) {
	logErr(logger.Fatal(), err)
}

// Fatalf logs a formatted message in fatal level and terminates the program.
func Fatalf(format string, args ...interface{}) {
	logMsg(logger.Fatal(), format, args...)
}

func logMsg(ev *zerolog.Event, format string, args ...interface{}) {
	var fmtArgs []interface{}
	for _, arg := range args {
		if arg, ok := arg.(Field); ok {
			switch v := arg.Value.(type) {
			case bool:
				ev.Bool(arg.Key, v)
			case int:
				ev.Int(arg.Key, v)
			case int64:
				ev.Int64(arg.Key, v)
			case string:
				ev.Str(arg.Key, v)
			case float64:
				ev.Float64(arg.Key, v)
			case time.Time:
				ev.Time(arg.Key, v)
			case time.Duration:
				ev.Dur(arg.Key, v)
			default:
				if err, ok := v.(error); ok {
					ev.AnErr(arg.Key, err)
				} else {
					ev.Interface(arg.Key, v)
				}
			}
		} else {
			fmtArgs = append(fmtArgs, arg)
		}
	}
	ev.Caller(2).Msgf(format, fmtArgs...)
}

func logErr(ev *zerolog.Event, err error) {
	switch v := err.(type) {
	case *service.Error:
		if v.Cause != nil {
			ev.Err(v.Cause)
		}
		ev.Caller(2).Str("class", v.Class.String()).
			Str("service", v.Service).
			Bool("temp", v.IsTemp).
			Msg(v.Message)
	default:
		if cause := errors.Unwrap(err); cause != nil {
			ev.Err(err)
		}
		ev.Caller(2).Str("class", service.EUnknown.String()).
			Msg(err.Error())
	}
}
