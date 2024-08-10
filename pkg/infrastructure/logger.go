package infrastructure

import (
	"io"
	"log"
)

type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
	traceLogger *log.Logger
}

// NewLogger cria uma nova instância de Logger com os escritores fornecidos
func NewLogger(infoHandle io.Writer, errorHandle io.Writer) *Logger {
	return &Logger{
		infoLogger:  log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		debugLogger: log.New(infoHandle, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		traceLogger: log.New(infoHandle, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Info registra uma mensagem de nível INFO
func (l *Logger) Info(msg string) {
	l.infoLogger.Println(msg)
}

// Error registra uma mensagem de nível ERROR
func (l *Logger) Error(msg string) {
	l.errorLogger.Println(msg)
}

// Debug registra uma mensagem de nível DEBUG
func (l *Logger) Debug(msg string) {
	l.infoLogger.Println(msg)
}

// Trace registra uma mensagem de nível TRACE
func (l *Logger) Trace(msg string) {
	l.infoLogger.Println(msg)
}
