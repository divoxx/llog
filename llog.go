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
	Debug Level = iota
	Info
	Warning
	Error
)

var (
	// levelName stores one letter representations for each level.
	levelPrefix = [4]rune{'D', 'I', 'W', 'E'}
)

// Log is an instance of a Log with a specific level set.
type Log struct {
	sync.Mutex
	io.Writer
	level Level
}

// New creates a new instance of Log that will log to the provided io.Writer only if the method used
// for logging is enabled for the provided level.
func New(w io.Writer, l Level) Log {
	return Log{Writer: w, level: l}
}

// Debug writes to the Log only if the log's level is set to Debug.
func (l *Log) Debug(msg ...interface{}) error {
	return l.log(Debug, msg...)
}

// Debugf writes a message formatted (as fmt.Printf) to the Log only if the log's level is set to Debug.
func (l *Log) Debugf(fmtStr string, msg ...interface{}) error {
	return l.logf(Debug, fmtStr, msg...)
}

// Info writes to the Log only if the log's level is set to Info.
func (l *Log) Info(msg ...interface{}) error {
	return l.log(Info, msg...)
}

// Infof writes a message formatted (as fmt.Printf) to the Log only if the log's level is set to Info.
func (l *Log) Infof(fmtStr string, msg ...interface{}) error {
	return l.logf(Info, fmtStr, msg...)
}

// Warning writes to the Log only if the log's level is set to Warning.
func (l *Log) Warning(msg ...interface{}) error {
	return l.log(Warning, msg...)
}

// Warningf writes a message formatted (as fmt.Printf) to the Log only if the log's level is set to Warning.
func (l *Log) Warningf(fmtStr string, msg ...interface{}) error {
	return l.logf(Warning, fmtStr, msg...)
}

// Error writes to the Log only if the log's level is set to Error.
func (l *Log) Error(msg ...interface{}) error {
	return l.log(Error, msg...)
}

// Errorf writes a message formatted (as fmt.Printf) to the Log only if the log's level is set to Error.
func (l *Log) Errorf(fmtStr string, msg ...interface{}) error {
	return l.logf(Error, fmtStr, msg...)
}

// log is called by all the other leveled logging functions.
func (l *Log) log(level Level, msg ...interface{}) error {
	if l.level > level {
		return nil
	}

	l.Lock()
	_, err := fmt.Fprintln(l, msg...)
	l.Unlock()
	return err
}

// logf is called by all the other leveled formatted logging functions.
func (l *Log) logf(level Level, fmtStr string, msg ...interface{}) error {
	if l.level > level {
		return nil
	}

	fmtStr = fmtStr + "\n"

	l.Lock()
	_, err := fmt.Fprintf(l, fmtStr, msg...)
	l.Unlock()
	return err
}
