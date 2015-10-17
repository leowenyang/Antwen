package antLog

import "testing"

func TestLog(t *testing.T) {
	Info("this is Info")
	Debug("this is Debug")
	Warning("this is Warning")
	Error("this is Error")
}
