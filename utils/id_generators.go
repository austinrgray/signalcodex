package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

const (
	TraceIdLength     = 16
	SpanIdLength      = 8
	RequestIdLength   = 8
	VesselIdLength    = 12
	FallbackTraceId   = "Fallback_TraceId"
	FallbackSpanId    = "Fallback"
	FallbackRequestId = "Fallback"
	FallbackVesselId  = "Fallback_VID"
)

func GenerateTraceID() (string, error) {
	return generateHex(TraceIdLength)
}

func GenerateSpanID() (string, error) {
	return generateHex(SpanIdLength)
}

func GenerateRequestID() (string, error) {
	return generateHex(RequestIdLength)
}

func generateHex(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return FallbackSpanId, fmt.Errorf("failed to generate random bytes: %w", err)
	}
	return hex.EncodeToString(bytes), nil
}

func GenerateVesselID() (string, error) {
	id, err := generateUUID()
	if err != nil {
		return FallbackVesselId, err
	}
	plainId := strings.ReplaceAll(id.String(), "-", "")
	return plainId[:VesselIdLength], nil
}

func generateUUID() (uuid.UUID, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to generate uuid: %w", err)
	}
	return id, nil
}
