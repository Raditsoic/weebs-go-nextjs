package utils

import (
	"context"
	"time"
)

func GetContextWithTimeout() (context.Context, context.CancelFunc) {
	timeoutSeconds := 30
	return context.WithTimeout(context.Background(), time.Duration(timeoutSeconds)*time.Second)
}
