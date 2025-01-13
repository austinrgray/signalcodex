package models

import (
	"fmt"
	"time"

	"github.com/austinrgray/signalcodex/messages"
	"go.uber.org/zap"

	pb "google.golang.org/protobuf/proto"
	timepb "google.golang.org/protobuf/types/known/timestamppb"
)

type ErrorSalience int

const (
	Low ErrorSalience = iota
	Medium
	High
	Critical
	Fatal
)

type Error struct {
	Code        string
	Salience    ErrorSalience
	Description string
	Timestamp   time.Time
}

func NewError(code string, salience ErrorSalience, desc string) Error {
	return Error{
		Code:        code,
		Salience:    salience,
		Description: desc,
		Timestamp:   time.Now(),
	}
}

func ErrorFromProto(msg *messages.Error) Error {
	return Error{
		Code:        msg.Code,
		Salience:    ErrorSalience(msg.Salience),
		Description: msg.Description,
		Timestamp:   msg.Timestamp.AsTime(),
	}
}

func ZapFromProto(msg *messages.Error) []zap.Field {
	err := fmt.Errorf("%s: %s", msg.Code, msg.Description)
	return []zap.Field{
		zap.Error(err),
		zap.Int("salience", int(msg.Salience)),
		zap.Time("timestamp", msg.Timestamp.AsTime()),
	}
}

func (e *Error) ToProto() *messages.Error {
	return &messages.Error{
		Code:        e.Code,
		Salience:    messages.ErrorSalience(e.Salience),
		Description: e.Description,
		Timestamp:   timepb.New(e.Timestamp),
	}
}

func MarshalError(msg *messages.Error) ([]byte, error) {
	bytes, err := pb.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal error message: %+v", err)
	}
	return bytes, nil
}

func UnmarshalError(bytes []byte) (*messages.Error, error) {
	var msg messages.Error
	err := pb.Unmarshal(bytes, &msg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal bytes to error message: %+v", err)
	}
	return &msg, nil
}
