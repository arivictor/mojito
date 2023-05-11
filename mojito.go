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

type Logger struct {
	out    io.Writer
	mu     sync.Mutex
	Config config
}

func New() *Logger {
	return &Logger{
		out: os.Stdout,
		Config: config{
			severityKey:  defaultSeverityKey,
			minLevel:     defaultLevel,
			messageKey:   defaultMessageKey,
			timestampKey: defaultTimeKey,
			delimiter:    defaultDelimiter,
			timeFormat:   time.RFC3339,
		},
	}
}

func (l *Logger) Debug(message string, properties ...interface{}) {
	l.print(Info, message, properties)
}

func (l *Logger) Info(message string, properties ...interface{}) {
	l.print(Info, message, properties)
}

func (l *Logger) Warn(message string, properties ...interface{}) {
	l.print(Info, message, properties)
}

func (l *Logger) Error(message string, properties ...interface{}) {
	l.print(Info, message, properties)
}

func (l *Logger) Fatal(message string, properties ...interface{}) {
	l.print(Info, message, properties)
	os.Exit(1)
}

func (l *Logger) print(level Level, message string, properties []interface{}) (int, error) {

	// Exit if not above desired log level
	if level < l.Config.minLevel {
		return 0, nil
	}

	r := l.makeRecord(level, message, properties)

	var line []byte

	line, err := json.Marshal(r)
	if err != nil {
		line = []byte(fmt.Sprintf("%s: unable to marshal log message: %s", Error.String(), err.Error()))
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	return l.out.Write(append(line, l.Config.delimiter))
}

func (l *Logger) Write(message []byte) (n int, err error) {
	return l.print(Error, string(message), nil)
}

func (l *Logger) makeRecord(level Level, message string, properties []interface{}) map[string]interface{} {
	r := make(map[string]interface{})
	r[l.Config.severityKey] = level.String()
	r[l.Config.timestampKey] = time.Now().Format(l.Config.timeFormat)
	r[l.Config.messageKey] = message

	if level >= Error {
		r[defaultTraceKey] = string(debug.Stack())
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
