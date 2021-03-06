// Copyright 2015 Antwen Author. All Rights Reserved.
//
// Usage:
//
// import "github.com/leowenyang/Antwen/antLog"
//
// antLog.Info("test")
// antLog.Debug("test")
// antLog.Warning("test")
// antLog.Error("test")
//

package antLog

import "log"

// output Info information
func Info(v ...interface{}) {
	log.Printf("[I]--> %v\n", v...)
}

// output Debug information
func Debug(v ...interface{}) {
	log.Printf("[D]--> %v\n", v...)
}

// output Warning information
func Warning(v ...interface{}) {
	log.Printf("[W]--> %v\n", v...)
}

// output Error information
func Error(v ...interface{}) {
	log.Fatalf("[E]--> %v\n", v...)
}
