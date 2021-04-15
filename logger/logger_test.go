package logger

import (
	"testing"
)

func TestInfo(t *testing.T) {
	l := Logger{}
	l.Info("info")
	l.Info("%s", "info")
}

func TestInfoPart(t *testing.T) {
	l := Logger{}
	l.InfoPart("info")
	l.InfoPart("%s", "info")
}

func TestFocused(t *testing.T) {
	l := Logger{}
	l.Focused("focused")
	l.Focused("%s", "focused")
}

func TestError(t *testing.T) {
	l := Logger{}
	l.Error("error")
	l.Error("%s", "error")
}

func TestErrors(t *testing.T) {
	l := Logger{}
	l.Errors([]string{
		"error1",
		"error2",
	})
	l.Errors([]string{
		"error%s",
		"error%s",
	}, "1", "2")
}

func TestListItem(t *testing.T) {
	l := Logger{}
	l.ListItem("item")
	l.ListItem("%s", "item")
}

func LogPart(t *testing.T) {
	l := Logger{}
	l.LogPart("message", StyleInfo)
	l.LogPart("%s", StyleInfo, "message")
}
