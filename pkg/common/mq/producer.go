package mq

import (
	"os"

	"github.com/apache/rocketmq-clients/golang/v5"
)

func init() {
	os.Setenv("mq.consoleAppender.enabled", "true")

	golang.ResetLogger()
}

func NewProducer(endpoint string, topic string) (golang.Producer, error) {
	// nil
	producer, err := golang.NewProducer(&golang.Config{
		Endpoint: endpoint,
	}, golang.WithTopics(topic))
	if err != nil {
		return nil, err
	}
	return producer, nil
}
