package logging

import (
	"github.com/austinrgray/signalcodex/models"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

func LoggerWithHeader(header nats.Header, logger *zap.Logger) *zap.Logger {
	fields := make([]zap.Field, 0)
	if v := header.Get(string(models.MsgVersionHeaderField)); v != "" {
		fields = append(fields, zap.String(string(models.MsgVersionHeaderField), v))
	}
	if v := header.Get(string(models.MsgFormatHeaderField)); v != "" {
		fields = append(fields, zap.String(string(models.MsgFormatHeaderField), v))
	}
	if v := header.Get(string(models.TimeStampHeaderField)); v != "" {
		fields = append(fields, zap.String(string(models.TimeStampHeaderField), v))
	}
	if v := header.Get(string(models.RetryCountHeaderField)); v != "" {
		fields = append(fields, zap.String(string(models.RetryCountHeaderField), v))
	}
	if v := header.Get(string(models.RequestIdHeaderField)); v != "" {
		fields = append(fields, zap.String(string(models.RequestIdHeaderField), v))
	}
	if v := header.Get(string(models.SpanIdHeaderField)); v != "" {
		fields = append(fields, zap.String(string(models.SpanIdHeaderField), v))
	}
	if v := header.Get(string(models.TraceIdHeaderField)); v != "" {
		fields = append(fields, zap.String(string(models.TraceIdHeaderField), v))
	}
	if v := header.Get(string(models.MsgTypeHeaderField)); v != "" {
		fields = append(fields, zap.String(string(models.MsgTypeHeaderField), v))
	}
	if v := header.Get(string(models.MsgSubtypeHeaderField)); v != "" {
		fields = append(fields, zap.String(string(models.MsgSubtypeHeaderField), v))
	}
	if v := header.Get(string(models.VesselIdHeaderField)); v != "" {
		fields = append(fields, zap.String(string(models.VesselIdHeaderField), v))
	}
	return logger.With(fields...)
}
