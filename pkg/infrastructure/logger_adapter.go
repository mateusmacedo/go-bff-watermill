package infrastructure

import (
	"fmt"
	"strings"

	"github.com/ThreeDotsLabs/watermill"
)

type WatermillLoggerAdapter struct {
	logger *Logger
}

func NewLoggerAdapter(logger *Logger) *WatermillLoggerAdapter {
	return &WatermillLoggerAdapter{logger: logger}
}

func (l *WatermillLoggerAdapter) Trace(msg string, fields watermill.LogFields) {
	l.logger.Info("[TRACE] " + msg + " " + formatFields(fields))
}

func (l *WatermillLoggerAdapter) Debug(msg string, fields watermill.LogFields) {
	l.logger.Info("[DEBUG] " + msg + " " + formatFields(fields))
}

func (l *WatermillLoggerAdapter) Info(msg string, fields watermill.LogFields) {
	l.logger.Info("[INFO] " + msg + " " + formatFields(fields))
}

func (l *WatermillLoggerAdapter) Error(msg string, err error, fields watermill.LogFields) {
	l.logger.Error("[ERROR] " + msg + ": " + err.Error() + " " + formatFields(fields))
}

func (l *WatermillLoggerAdapter) With(fields watermill.LogFields) watermill.LoggerAdapter {
	return &WatermillLoggerAdapter{
		logger: l.logger,
	}
}

// formatFields converte o mapa de fields em uma string legível
func formatFields(fields watermill.LogFields) string {
	var sb strings.Builder
	for key, value := range fields {
		sb.WriteString(fmt.Sprintf("%s=%v ", key, value))
	}
	return sb.String()
}
