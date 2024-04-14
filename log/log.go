package log

import "log"

var Log *log.Logger

func LogInit() {
	Log = log.Default()
}
