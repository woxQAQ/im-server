package utils

import (
	"fmt"
	"time"
)

type Retry interface {
	Do(Operation) error
}

type Operation func() error

type SimpleRetry struct {
	maxRetries int
	sleepTime  time.Duration
}

func NewSimpleRetry(maxRetries int, sleepTime time.Duration) SimpleRetry {
	return SimpleRetry{
		maxRetries: maxRetries,
		sleepTime:  sleepTime,
	}
}

func (s SimpleRetry) Do(callback Operation) error {
	var err error
	for i := 0; i < s.maxRetries; i++ {
		err = callback()
		if err == nil {
			return nil
		}

		fmt.Printf("Operation failed (attempt %d): %v\n", i+1, err)
		time.Sleep(s.sleepTime)
	}

	return fmt.Errorf("Operation failed due to %v\n", err)
}
