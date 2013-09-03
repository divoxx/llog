/*
llog is a simple and straightforward implementation of leveled logs for go.

The levels are defined as follow: Debug < Info < Warning < Error.

A call to any of the logging methods will only append to the log if the level
represented by the method is higher or equal to the level set in the Log instance.

For example: a Log w/ level Warning will only append to the log any call to Error,
Errorf, Warning and Warningf. All other calls will be ignored.

	l := llog.New(os.Stdout, llog.Warning)
	l.Warning("Something wrong occurred")	// Gets logged
	l.Info("You might like to know this")	// Doesn't get logged

*/
package llog

import (
	"fmt"
	"io"
	"sync"
)

// Level represents log levels which are sorted based on the underlying number associated to it.
type Level uint8

// Definitions of available levels: Debug < Info < Warning < Error.
const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
)

var (
	// levelName stores one letter representations for each level.
	levelPrefix = [4]rune{'D', 'I', 'W', 'E'}
)

// Log is an instance that contains a reference to an io.Writer and a level set. For more info,
// check New documentation.
type Log struct {
	m     sync.Mutex
	w     io.Writer
	level Level
}

// New creates a new instance of Log that will log to the provided io.Writer only if the method used
// for logging is enabled for the provided level. See package documentation for more details and examples.
func New(w io.Writer, l Level) Log {
	return Log{w: w, level: l}
}

// SetLevel changes the logging level for the log instance.
func (l *Log) SetLeveL(l Level) {
	l.level = Level
}

// Debug writes to the Log only if the log's level is set to Debug.
func (l *Log) Debug(msg ...interface{}) error {
	return l.log(DEBUG, msg...)
}

// Debugf writes a message formatted (as fmt.Printf) to the Log only if the log's level is set to Debug.
func (l *Log) Debugf(fmtStr string, msg ...interface{}) error {
	return l.logf(DEBUG, fmtStr, msg...)
}

// Info writes to the Log only if the log's level is set to Info.
func (l *Log) Info(msg ...interface{}) error {
	return l.log(INFO, msg...)
}

// Infof writes a message formatted (as fmt.Printf) to the Log only if the log's level is set to Info.
func (l *Log) Infof(fmtStr string, msg ...interface{}) error {
	return l.logf(INFO, fmtStr, msg...)
}

// Warning writes to the Log only if the log's level is set to Warning.
func (l *Log) Warning(msg ...interface{}) error {
	return l.log(WARNING, msg...)
}

// Warningf writes a message formatted (as fmt.Printf) to the Log only if the log's level is set to Warning.
func (l *Log) Warningf(fmtStr string, msg ...interface{}) error {
	return l.logf(WARNING, fmtStr, msg...)
}

// Error writes to the Log only if the log's level is set to Error.
func (l *Log) Error(msg ...interface{}) error {
	return l.log(ERROR, msg...)
}

// Errorf writes a message formatted (as fmt.Printf) to the Log only if the log's level is set to Error.
func (l *Log) Errorf(fmtStr string, msg ...interface{}) error {
	return l.logf(ERROR, fmtStr, msg...)
}

// log is called by all the other leveled logging functions.
func (l *Log) log(level Level, msg ...interface{}) error {
	if l.level > level {
		return nil
	}

	l.m.Lock()
	_, err := fmt.Fprintln(l.w, msg...)
	l.m.Unlock()
	return err
}

// logf is called by all the other leveled formatted logging functions.
func (l *Log) logf(level Level, fmtStr string, msg ...interface{}) error {
	if l.level > level {
		return nil
	}

	fmtStr = fmtStr + "\n"

	l.m.Lock()
	_, err := fmt.Fprintf(l.w, fmtStr, msg...)
	l.m.Unlock()
	return err
}
