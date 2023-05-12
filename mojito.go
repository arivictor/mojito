package mojito

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

const (
	DEFAULT_SEVERITY_KEY = "severity"
	DEFAULT_TIME_KEY     = "time"
	DEFAULT_MESSAGE_KEY  = "message"
	DEFAULT_TRACE_KEY    = "trace"
	DEFAULT_DELIMITER    = '\n'
	DEFAULT_MIN_LEVEL    = INFO
	DEFAULT_TIME_FORMAT  = time.RFC3339
)

type Logger struct {
	out    io.Writer
	mu     sync.Mutex
	Config config
}

func New() *Logger {
	c := config{
		severityKey:  DEFAULT_SEVERITY_KEY,
		minLevel:     DEFAULT_MIN_LEVEL,
		messageKey:   DEFAULT_MESSAGE_KEY,
		timestampKey: DEFAULT_TIME_KEY,
		delimiter:    DEFAULT_DELIMITER,
		timeFormat:   DEFAULT_TIME_FORMAT,
	}

	return &Logger{
		out:    os.Stdout,
		Config: c,
	}
}

func (l *Logger) Debug(message string, properties ...interface{}) {
	l.print(DEBUG, message, properties)
}

func (l *Logger) Info(message string, properties ...interface{}) {
	l.print(INFO, message, properties)
}

func (l *Logger) Warn(message string, properties ...interface{}) {
	l.print(WARN, message, properties)
}

func (l *Logger) Error(message string, properties ...interface{}) {
	l.print(ERROR, message, properties)
}

func (l *Logger) Fatal(message string, properties ...interface{}) {
	l.print(FATAL, message, properties)
	os.Exit(1)
}

func (l *Logger) print(level Level, message string, properties []interface{}) (int, error) {

	if level < l.Config.minLevel {
		return 0, nil
	}

	r := l.makeRecord(level, message, properties)

	var line []byte
	line, err := json.Marshal(r)
	if err != nil {
		line = []byte(fmt.Sprintf("%s: unable to marshal log message: %s", ERROR.String(), err.Error()))
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	return l.out.Write(append(line, l.Config.delimiter))
}

func (l *Logger) Write(message []byte) (n int, err error) {
	return l.print(ERROR, string(message), nil)
}

func (l *Logger) makeRecord(level Level, message string, properties []interface{}) map[string]interface{} {
	r := make(map[string]interface{})
	r[l.Config.severityKey] = level.String()
	r[l.Config.timestampKey] = time.Now().Format(l.Config.timeFormat)
	r[l.Config.messageKey] = message

	if level >= ERROR {
		r[DEFAULT_TRACE_KEY] = string(debug.Stack())
	}

	for i, prop := range properties {
		if i+1 == len(properties) {
			r[fmt.Sprintf("%v", prop)] = ""
		} else {
			r[fmt.Sprintf("%v", prop)] = properties[i+1]
		}
	}

	return r
}
