package logger

import "log"

// Default log level is debug 1000
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

// Error : method for print error logger
func Error(data interface{}) {
	if IsErrorEnabled() {
		log.Println(data)
	}
}

// IsErrorEnabled : method for check error enabled
func IsErrorEnabled() bool {
	return (logLevel & 0x01) > 0
}

// Fatal : method for print fatal logger
func Fatal(data interface{}) {
	log.Fatal(data)
}
