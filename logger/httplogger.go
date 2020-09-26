package logger

import "log"

// Default log level is debug
var logLevel uint8 = 0x08

// Debug : method for print Debug logger
func Debug(data interface{}) {
	if IsDebugEnabled() {
		log.Println(data)
	}
}

// IsDebugEnabled : method for check debug enabled
func IsDebugEnabled() bool {
	return (logLevel & 0x08) > 0
}
