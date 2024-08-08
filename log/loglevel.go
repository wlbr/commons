package log

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
// Level DEBUG should be used for temporary debugging information and should be removed
// from production code.
type LogLevel int

// The predefined LogLevels that are used by the logging funktions below.
const (
	OFF LogLevel = iota
	FATAL
	ERROR
	WARN
	INFO
	DEBUG
	ALL
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
	} else if strings.ToUpper(logfilename) == "STDOUT" {
		lfilename = "<STDOUT>"
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
	case FATAL:
		c = color.Red
	case ERROR:
		c = color.LightRed
	case WARN:
		c = color.Yellow
	case DEBUG:
		c = color.LightBlue
	case INFO:
		c = color.LightCyan
	}
	return c.Sprintf("%s", s)
}

func (l *Logger) writelog(level LogLevel, format string, args ...interface{}) {
	if level == DEBUG || l.ActiveLoglevel >= level {
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
	l.writelog(INFO, format, args...)
}

// Warn works just as fmt.Printf, but prints into the loggers stream.
// The message is only printed if ActiveLogLevel is set higher or equal to 'Warn'
func (l *Logger) Warn(format string, args ...interface{}) {
	l.writelog(WARN, format, args...)
}

// Error works just as fmt.Printf, but prints into the loggers stream.
// The message is only printed if ActiveLogLevel is set higher or equal to 'Error'
func (l *Logger) Error(format string, args ...interface{}) {
	l.writelog(ERROR, format, args...)
}

// Fatal works just as fmt.Printf, but prints into the loggers stream.
// The message is only printed if ActiveLogLevel is set hogher or equal to 'Fatal'
func (l *Logger) Fatal(format string, args ...interface{}) {
	l.writelog(FATAL, format, args...)
}

// Debug works just as fmt.Printf, but prints into the loggers stream.
// The message is printed anyway, regardless of the log level.
func (l *Logger) Debug(format string, args ...interface{}) {
	l.writelog(DEBUG, format, args...)
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
		convenienceLogger.writelog(INFO, format, args...)
	} else {
		outputToStandardLogger(INFO, format, args...)
	}
}

// LogDebug works just as fmt.Printf, but prints into the Convenience loggers stream, as set with
// SetConvenienceLogger(). It uses the standard logger (package log) if te Convenience logger is unset.
// The message is only printed if ActiveLogLevel is set higher or equal to 'Debug'
func Debug(format string, args ...interface{}) {
	if convenienceLogger != nil {
		convenienceLogger.writelog(DEBUG, format, args...)
	} else {
		outputToStandardLogger(DEBUG, format, args...)
	}
}

// LogWarn works just as fmt.Printf, but prints into the Convenience loggers stream, as set with
// SetConvenienceLogger(). It uses the standard logger (package log) if te Convenience logger is unset.
// The message is only printed if ActiveLogLevel is set higher or equal to 'Warn'
func Warn(format string, args ...interface{}) {
	if convenienceLogger != nil {
		convenienceLogger.writelog(WARN, format, args...)
	} else {
		outputToStandardLogger(WARN, format, args...)
	}
}

// LogError works just as fmt.Printf, but prints into the Convenience loggers stream, as set with
// SetConvenienceLogger(). It uses the standard logger (package log) if te Convenience logger is unset.
// The message is only printed if ActiveLogLevel is set higher or equal to 'Error'
func Error(format string, args ...interface{}) {
	if convenienceLogger != nil {
		convenienceLogger.writelog(ERROR, format, args...)
	} else {
		outputToStandardLogger(ERROR, format, args...)
	}
}

// Fatal works just as fmt.Printf, but prints into the Convenience loggers stream, as set with
// SetConvenienceLogger(). It uses the standard logger (package log) if te Convenience logger is unset.
// The message is only printed if ActiveLogLevel is set hogher or equal to 'Fatal'
func Fatal(format string, args ...interface{}) {
	if convenienceLogger != nil {
		convenienceLogger.writelog(FATAL, format, args...)
	} else {
		outputToStandardLogger(FATAL, format, args...)
	}
}
