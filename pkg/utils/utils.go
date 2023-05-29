package utils

import (
	"time"

	"github.com/wispeeer/wispeeer/pkg/logger"
)

func Timer(action string, start time.Time) {
	dis := time.Since(start)
	if dis > 1 {
		logger.Task(action).Printf("Task done in %v", dis)
	}
}
