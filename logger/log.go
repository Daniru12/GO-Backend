package logger

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"

	. "github.com/logrusorgru/aurora"
)

var nativeLog *log.Logger
var errorLog *log.Logger

var (
	fatal = `FATAL`
	err   = `ERR`
	warn  = `WARN`
	info  = `INFO`
	debug = `DEBUG`
	trace = `TRACE`
)

var (
	fatalColor = BgRed(`[FATAL]`)
	errorColor = BgRed(`[ERROR]`)
	warnColor  = BgBrown(`[WARN]`)
	infoColor  = BgBlue(`[INFO]`)
	debugColor = BgCyan(`[DEBUG]`)
	traceColor = BgMagenta(`[TRACE]`)
)

var logTypes = map[string]int{
	`FATAL`: 1,
	`ERROR`: 2,
	`WARN`:  3,
	`INFO`:  4,
	`DEBUG`: 5,
	`TRACE`: 6,
}

func init() {
	nativeLog = log.New(os.Stdout, ``, log.LstdFlags|log.Lmicroseconds)
	errorLog = log.New(os.Stderr, ``, log.LstdFlags|log.Lmicroseconds)
}

// isLoggable Check whether the log type is loggable under current configurations
func isLoggable(logType string) bool {
	return logTypes[logType] <= logTypes[logConfig.Level]
}

func toString(id string, prefix Value, message interface{}, params interface{}, file string, line int) string {

	var messageFmt = "%s %s, [%v]"
	var paramsFormatted = fmt.Sprint(params)
	if len(paramsFormatted) == 11 {
		paramsFormatted = ``
	}

	if logConfig.FilePath {
		messageFmt = "%s %s, %v On %s at line %d"
		return fmt.Sprintf(messageFmt,
			prefix,
			fmt.Sprintf(`%+v`, message),
			paramsFormatted,
			file,
			line)
	}

	return fmt.Sprintf(messageFmt,
		prefix,
		fmt.Sprintf(`%+v`, message),
		paramsFormatted)
}

func ErrorContext(ctx context.Context, message interface{}, params ...interface{}) {
	logEntryContext(err, ctx, message, errorColor, params)
}

func WarnContext(ctx context.Context, message interface{}, params ...interface{}) {
	logEntryContext(warn, ctx, message, warnColor, params)
}

func InfoContext(ctx context.Context, message interface{}, params ...interface{}) {
	logEntryContext(info, ctx, message, infoColor, params)
}

func DebugContext(ctx context.Context, message interface{}, params ...interface{}) {
	logEntryContext(debug, ctx, message, debugColor, params)
}

func TraceContext(ctx context.Context, message interface{}, params ...interface{}) {
	logEntryContext(trace, ctx, message, traceColor, params)
}

func Error(message interface{}, params ...interface{}) {
	logEntry(err, message, errorColor, params)
}

func Warn(message interface{}, params ...interface{}) {
	logEntry(warn, message, warnColor, params)
}

func Info(message interface{}, params ...interface{}) {
	logEntry(info, message, infoColor, params)
}

func Debug(message interface{}, params ...interface{}) {
	logEntry(debug, message, debugColor, params)
}

func Trace(message interface{}, params ...interface{}) {
	logEntry(trace, message, traceColor, params)
}

func Fatal(message interface{}, params ...interface{}) {
	logEntry(fatal, message, fatalColor, params)
}

func Fataln(message interface{}, params ...interface{}) {
	logEntry(fatal, message, fatalColor, params)
}

func FatalContext(ctx context.Context, message interface{}, params interface{}) {
	logEntry(fatal, message, fatalColor, params)
}

func logEntryContext(logType string, ctx context.Context, message interface{}, color Value, params interface{}) {

	if !isLoggable(logType) {
		return
	}

	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = `<Unknown>`
		line = 1
	}

	if logType == fatal {
		errorLog.Fatalln(toString(``, color, message, params, file, line))
	}

	if logType == err {
		errorLog.Println(toString(``, color, message, params, file, line))
		return
	}

	nativeLog.Println(toString(``, color, message, params, file, line))
}

func WithPrefix(p string, message interface{}) string {
	return fmt.Sprintf(`[%s] [%+v]`, p, message)
}

func logEntry(logType string, message interface{}, color Value, params interface{}) {

	if !isLoggable(logType) {
		return
	}

	var file string
	var line int
	if logConfig.FilePath {
		_, f, l, ok := runtime.Caller(2)
		if !ok {
			f = `<Unknown>`
			l = 1
		}

		file = f
		line = l
	}

	if logType == fatal {
		nativeLog.Fatalln(toString(``, color, message, params, file, line))
	}

	if logType == err {
		nativeLog.Println(toString(``, color, message, params, file, line))
		return
	}

	nativeLog.Println(toString(``, color, message, params, file, line))
}
