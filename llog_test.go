package llog

import (
	"bytes"
	"fmt"
	"testing"
)

func TestOnlyErrorLogging(t *testing.T) {
	b := new(bytes.Buffer)
	l := New(b, ERROR)

	logAndAssertNoError(t, l.Error, "Error")
	logfAndAssertNoError(t, l.Errorf, "%s", "Errorf")

	logAndAssertNoError(t, l.Warning, "Warning")
	logfAndAssertNoError(t, l.Warningf, "%s", "Warningf")

	logAndAssertNoError(t, l.Info, "Info")
	logfAndAssertNoError(t, l.Infof, "%s", "Infof")

	logAndAssertNoError(t, l.Debug, "Debug")
	logfAndAssertNoError(t, l.Debugf, "%s", "Debugf")

	verifyLogEntries(t, b, "Error", "Errorf")
}

func TestWarningOrHigher(t *testing.T) {
	b := new(bytes.Buffer)
	l := New(b, WARNING)

	logAndAssertNoError(t, l.Error, "Error")
	logfAndAssertNoError(t, l.Errorf, "%s", "Errorf")

	logAndAssertNoError(t, l.Warning, "Warning")
	logfAndAssertNoError(t, l.Warningf, "%s", "Warningf")

	logAndAssertNoError(t, l.Info, "Info")
	logfAndAssertNoError(t, l.Infof, "%s", "Infof")

	logAndAssertNoError(t, l.Debug, "Debug")
	logfAndAssertNoError(t, l.Debugf, "%s", "Debugf")

	verifyLogEntries(t, b, "Error", "Errorf", "Warning", "Warningf")
}

func TestInfoOrHigher(t *testing.T) {
	b := new(bytes.Buffer)
	l := New(b, INFO)

	logAndAssertNoError(t, l.Error, "Error")
	logfAndAssertNoError(t, l.Errorf, "%s", "Errorf")

	logAndAssertNoError(t, l.Warning, "Warning")
	logfAndAssertNoError(t, l.Warningf, "%s", "Warningf")

	logAndAssertNoError(t, l.Info, "Info")
	logfAndAssertNoError(t, l.Infof, "%s", "Infof")

	logAndAssertNoError(t, l.Debug, "Debug")
	logfAndAssertNoError(t, l.Debugf, "%s", "Debugf")

	verifyLogEntries(t, b, "Error", "Errorf", "Warning", "Warningf", "Info", "Infof")
}

func TestDebugOrHigher(t *testing.T) {
	b := new(bytes.Buffer)
	l := New(b, DEBUG)

	logAndAssertNoError(t, l.Error, "Error")
	logfAndAssertNoError(t, l.Errorf, "%s", "Errorf")

	logAndAssertNoError(t, l.Warning, "Warning")
	logfAndAssertNoError(t, l.Warningf, "%s", "Warningf")

	logAndAssertNoError(t, l.Info, "Info")
	logfAndAssertNoError(t, l.Infof, "%s", "Infof")

	logAndAssertNoError(t, l.Debug, "Debug")
	logfAndAssertNoError(t, l.Debugf, "%s", "Debugf")

	verifyLogEntries(t, b, "Error", "Errorf", "Warning", "Warningf", "Info", "Infof", "Debug", "Debugf")
}

func logAndAssertNoError(t *testing.T, fn func(...interface{}) error, msg ...interface{}) {
	err := fn(msg...)
	if err != nil {
		t.Fatal(err)
	}
}

func logfAndAssertNoError(t *testing.T, fn func(string, ...interface{}) error, fmt string, msg ...interface{}) {
	err := fn(fmt, msg...)
	if err != nil {
		t.Fatal(err)
	}
}

func verifyLogEntries(t *testing.T, b *bytes.Buffer, entries ...string) {
	eb := new(bytes.Buffer)

	for _, e := range entries {
		eb.WriteString(fmt.Sprintln(e))
	}

	if eb.String() != b.String() {
		t.Errorf("Expected entries log does not match\nExpected: %s\nGot: %s\n", eb.String(), b.String())
	}
}
