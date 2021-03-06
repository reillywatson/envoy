package messenger

import "log"

type Logger interface {
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Panic(err interface{}, stack string)
}

type defaultLogger struct{}

func (l defaultLogger) Debugf(format string, v ...interface{}) {
	log.Printf("DEBUG: "+format, v...)
}

func (l defaultLogger) Infof(format string, v ...interface{}) {
	log.Printf("INFO:  "+format, v...)
}

func (l defaultLogger) Errorf(format string, v ...interface{}) {
	log.Printf("ERROR: "+format, v...)
}

func (l defaultLogger) Panic(err interface{}, stack string) {
	log.Printf("PANIC: %v\nStack:\n%s", err, stack)
}
