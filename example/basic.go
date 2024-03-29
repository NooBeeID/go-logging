package example

import (
	"context"

	"github.com/NooBeeID/go-logging/logger"
	"github.com/sirupsen/logrus"
)

func Basic() {
	log := logger.NewLog()

	log.SetLevel(int(logrus.DebugLevel))

	ctx := context.Background()
	log.Infof(ctx, "finish")
	log.Errorf(ctx, "error")
}
