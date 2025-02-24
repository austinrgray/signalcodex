package logging

import (
	"context"
	"fmt"

	"github.com/austinrgray/signalcodex/messages"
	"go.uber.org/zap"
)

func LogSystemAlert(ctx context.Context, logger *zap.Logger, msg *messages.SystemAlert) {

}

func ZapSystemAlert(msg *messages.SystemAlert) []zap.Field {
	err := fmt.Errorf("%s : %s", msg.Code, msg.Description)
	return []zap.Field{
		zap.Error(err),
		zap.Int(string(LogKeySalience), int(msg.Salience)),
		zap.Time(string(LogKeyTimeStamp), msg.Timestamp.AsTime()),
	}
}
