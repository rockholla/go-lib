// Package logger is a general-use CLI tool logger
package logger

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

var (
	// StyleError represents the log item style for an error
	StyleError = color.New(color.FgHiRed, color.Bold)
	// StyleFocused represents a log item that should be able to picked out of the mix
	StyleFocused = color.New(color.Bold)
	// StyleInfo is style for an info, or normal, message
	StyleInfo = color.New()
	s         *spinner.Spinner
)

// Interface is an interface for a cli logger
type Interface interface {
	Focused(message string, args ...interface{})
	Error(message string, args ...interface{})
	Errors(messages []string, args ...interface{})
	Info(message string, args ...interface{})
	InfoPart(message string, args ...interface{})
	ListItem(message string, args ...interface{})
	LogPart(message string, style *color.Color, args ...interface{})
	Log(message string, style *color.Color, args ...interface{})
	SpinnerStart(message string)
	SpinnerStop()
}

// Logger is the default logger for vshasta
type Logger struct{}

// Info will log a normal info message
func (l *Logger) Info(message string, args ...interface{}) {
	l.Log(message, StyleInfo, args...)
}

// InfoPart will log part of line, without breaking so that other log items can
// still be appended to the line
func (l *Logger) InfoPart(message string, args ...interface{}) {
	l.LogPart(message, StyleInfo, args...)
}

// Focused will log a message that should stand out in the mix
func (l *Logger) Focused(message string, args ...interface{}) {
	l.Log(fmt.Sprintf("\n%s\n", message), StyleFocused, args...)
}

// Error will log an error
func (l *Logger) Error(message string, args ...interface{}) {
	l.Log(message, StyleError, args...)
}

// Errors will log multiple errors
func (l *Logger) Errors(messages []string, args ...interface{}) {
	message := "Errors:\n"
	for _, err := range messages {
		message = fmt.Sprintf("%s • %s\n", message, err)
	}
	l.Error(message, args...)
}

// ListItem will log a line as an item in a list
func (l *Logger) ListItem(message string, args ...interface{}) {
	l.Info(fmt.Sprintf(" • %s", message), args...)
}

// LogPart will log part of line, without breaking so that other log items can
// still be appended to the line
func (l *Logger) LogPart(message string, style *color.Color, args ...interface{}) {
	printer := style.SprintfFunc()
	fmt.Print(printer(message, args...))
}

// Log is the main, generic method for outputting log messages
func (l *Logger) Log(message string, style *color.Color, args ...interface{}) {
	printer := style.SprintfFunc()
	fmt.Println(printer(message, args...))
}

// SpinnerStart will start a spinner/waiting message that will go until we stop it
func (l *Logger) SpinnerStart(message string) {
	fmt.Print(fmt.Sprintf("%s...", message))
	s = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Start()
}

// SpinnerStop will stop/cancel a spinner
func (l *Logger) SpinnerStop() {
	if s.Active() {
		s.Stop()
	}
}
