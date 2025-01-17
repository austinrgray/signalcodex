package models

import (
	"fmt"
	"time"

	"github.com/austinrgray/signalcodex/messages"

	pb "google.golang.org/protobuf/proto"
	timepb "google.golang.org/protobuf/types/known/timestamppb"
)

type SystemAlertSalience int

const (
	Low SystemAlertSalience = iota
	Medium
	High
	Critical
	Fatal
)

type SystemAlert struct {
	Code        string
	Salience    SystemAlertSalience
	Description string
	Timestamp   time.Time
}

func NewSystemAlert(code string, salience SystemAlertSalience, desc string) SystemAlert {
	return SystemAlert{
		Code:        code,
		Salience:    salience,
		Description: desc,
		Timestamp:   time.Now(),
	}
}

func SystemAlertFromProto(msg *messages.SystemAlert) SystemAlert {
	return SystemAlert{
		Code:        msg.Code,
		Salience:    SystemAlertSalience(msg.Salience),
		Description: msg.Description,
		Timestamp:   msg.Timestamp.AsTime(),
	}
}

func (e *SystemAlert) ToProto() *messages.SystemAlert {
	return &messages.SystemAlert{
		Code:        e.Code,
		Salience:    messages.SystemAlertSalience(e.Salience),
		Description: e.Description,
		Timestamp:   timepb.New(e.Timestamp),
	}
}

func MarshalError(msg *messages.SystemAlert) ([]byte, error) {
	bytes, err := pb.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal error message: %+v", err)
	}
	return bytes, nil
}

func UnmarshalError(bytes []byte) (*messages.SystemAlert, error) {
	var msg messages.SystemAlert
	err := pb.Unmarshal(bytes, &msg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal bytes to error message: %+v", err)
	}
	return &msg, nil
}
