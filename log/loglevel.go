package logger

//go:generate enumer -type LogLevel loglevel.go

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/gookit/color"
)

// LogLevel sets the criticality of a logging output. It is used to filter logging messages
// depending on their priority. Compare to log4j.
type LogLevel int

// The predefined LogLevels that are used by the logging funktions below.
const (
	Off LogLevel = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	DebugLevel
	InfoLevel
	All
)

var loggerflags = log.Ldate | log.Ltime | log.Llongfile | log.Lmicroseconds | log.LUTC
var convenienceLogger *Logger
var colorizedOutput bool = true

// A Logger is an onbject the offers several method to write Messages to a stream.
// Atually it is a wrapper aroung the 'log' package, that enhances the LogLevel functionality.
type Logger struct {
	internallogger    *log.Logger
	ActiveLoglevel    LogLevel
	UseColouredOutput bool
}

// NewLoggerFromFile creates a new Logger. It take a file parameter (io.Writer) output file
// and a LogLevel to filter the messages that are wanted.
// The first created logger wil be set to be the convenience logger (see the convenience
// functions Log...() below). Afterwards created loggers will not overwrie this. The convenience
// logger can be reset by using the method SetConvenienceLogger.
func NewLoggerFromFile(logfile io.Writer, level LogLevel, useColouredOutput bool) *Logger {
	l := &Logger{}
	l.internallogger = log.New(logfile, "LOG: ", loggerflags)
	l.ActiveLoglevel = level
	l.UseColouredOutput = useColouredOutput
	if convenienceLogger == nil {
		convenienceLogger = l
	}
	return l
}

// NewLogger creates a new Logger. It take a string file name as output file
// and a LogLevel to filter the messages that are wanted.
// The logger will use io.StdOut if the log filename string parameter is "STDERR"
func NewLogger(logfilename string, level LogLevel, useColouredLogging ...bool) *Logger {
	var lfilename string
	var logfile io.Writer
	var useColouredOutput bool = false
	if len(useColouredLogging) > 0 {
		useColouredOutput = useColouredLogging[0]
	}
	if logfilename == "" || strings.ToUpper(logfilename) == "STDERR" {
		lfilename = "<STDERR>"
		logfile = os.Stderr
		if len(useColouredLogging) == 0 {
			useColouredOutput = true
		}
	} else {
		lfilename = logfilename
		logfile, _ = os.OpenFile(lfilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	}
	return NewLoggerFromFile(logfile, level, useColouredOutput)
}

func colorize(level LogLevel, s string) string {
	var c color.Color
	switch level {
	case FatalLevel:
		c = color.Red
	case ErrorLevel:
		c = color.LightRed
	case WarnLevel:
		c = color.Yellow
	case DebugLevel:
		c = color.LightBlue
	case InfoLevel:
		c = color.LightCyan
	}
	return c.Sprintf("%s", s)
}

func (l *Logger) writelog(level LogLevel, format string, args ...interface{}) {
	if l.ActiveLoglevel >= level {
		prefix := fmt.Sprintf("%5s: ", strings.TrimSuffix(strings.ToUpper(level.String()), "LEVEL"))
		if l.UseColouredOutput {
			prefix = colorize(level, prefix)
		}
		l.internallogger.SetPrefix(prefix)
		l.internallogger.Output(3, fmt.Sprintf(format, args...))
	}
}

// Info works just as fmt.Printf, but prints into the loggers stream.
// The message is only printed if ActiveLogLevel is set higher or equal to 'Info'
func (l *Logger) Info(format string, args ...interface{}) {
	l.writelog(InfoLevel, format, args...)
}

// Debug works just as fmt.Printf, but prints into the loggers stream.
// The message is only printed if ActiveLogLevel is set higher or equal to 'Debug'
func (l *Logger) Debug(format string, args ...interface{}) {
	l.writelog(DebugLevel, format, args...)
}

// Warn works just as fmt.Printf, but prints into the loggers stream.
// The message is only printed if ActiveLogLevel is set higher or equal to 'Warn'
func (l *Logger) Warn(format string, args ...interface{}) {
	l.writelog(WarnLevel, format, args...)
}

// Error works just as fmt.Printf, but prints into the loggers stream.
// The message is only printed if ActiveLogLevel is set higher or equal to 'Error'
func (l *Logger) Error(format string, args ...interface{}) {
	l.writelog(ErrorLevel, format, args...)
}

// Fatal works just as fmt.Printf, but prints into the loggers stream.
// The message is only printed if ActiveLogLevel is set hogher or equal to 'Fatal'
func (l *Logger) Fatal(format string, args ...interface{}) {
	l.writelog(FatalLevel, format, args...)
}

// -----------------------------

// SetConvenienceLogger sets a logger as a singleton object. The LogInfo etc.
// functions use this singleton to offer logging function without an object context.
func (l *Logger) SetConvenienceLogger() {
	convenienceLogger = l
}

func outputToStandardLogger(level LogLevel, format string, args ...interface{}) {
	p := log.Prefix()
	f := log.Flags()
	prefix := fmt.Sprintf("%5s: ", strings.TrimSuffix(strings.ToUpper(level.String()), "LEVEL"))
	if colorizedOutput {
		prefix = colorize(level, prefix)
	}
	log.SetPrefix(prefix)
	log.Output(3, fmt.Sprintf(format, args...))
	log.SetPrefix(p)
	log.SetFlags(f)
}

// LogInfo works just as fmt.Printf, but prints into the Convenience loggers stream, as set with
// SetConvenienceLogger(). It uses the standard logger (package log) if te Convenience logger is unset.
// The message is only printed if ActiveLogLevel is set higher or equal to 'Info'
func Info(format string, args ...interface{}) {
	if convenienceLogger != nil {
		convenienceLogger.writelog(InfoLevel, format, args...)
	} else {
		outputToStandardLogger(InfoLevel, format, args...)
	}
}

// LogDebug works just as fmt.Printf, but prints into the Convenience loggers stream, as set with
// SetConvenienceLogger(). It uses the standard logger (package log) if te Convenience logger is unset.
// The message is only printed if ActiveLogLevel is set higher or equal to 'Debug'
func Debug(format string, args ...interface{}) {
	if convenienceLogger != nil {
		convenienceLogger.writelog(DebugLevel, format, args...)
	} else {
		outputToStandardLogger(DebugLevel, format, args...)
	}
}

// LogWarn works just as fmt.Printf, but prints into the Convenience loggers stream, as set with
// SetConvenienceLogger(). It uses the standard logger (package log) if te Convenience logger is unset.
// The message is only printed if ActiveLogLevel is set higher or equal to 'Warn'
func Warn(format string, args ...interface{}) {
	if convenienceLogger != nil {
		convenienceLogger.writelog(WarnLevel, format, args...)
	} else {
		outputToStandardLogger(WarnLevel, format, args...)
	}
}

// LogError works just as fmt.Printf, but prints into the Convenience loggers stream, as set with
// SetConvenienceLogger(). It uses the standard logger (package log) if te Convenience logger is unset.
// The message is only printed if ActiveLogLevel is set higher or equal to 'Error'
func Error(format string, args ...interface{}) {
	if convenienceLogger != nil {
		convenienceLogger.writelog(ErrorLevel, format, args...)
	} else {
		outputToStandardLogger(ErrorLevel, format, args...)
	}
}

// Fatal works just as fmt.Printf, but prints into the Convenience loggers stream, as set with
// SetConvenienceLogger(). It uses the standard logger (package log) if te Convenience logger is unset.
// The message is only printed if ActiveLogLevel is set hogher or equal to 'Fatal'
func Fatal(format string, args ...interface{}) {
	if convenienceLogger != nil {
		convenienceLogger.writelog(FatalLevel, format, args...)
	} else {
		outputToStandardLogger(FatalLevel, format, args...)
	}
}
