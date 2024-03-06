package mq

import (
	"os"

	"github.com/apache/rocketmq-clients/golang/v5"
)

func init() {
	os.Setenv("mq.consoleAppender.enabled", "true")

	golang.ResetLogger()
}

// func NewProducer(endpoint string, topic string) (golang.Producer, error) {
// 	// nil
// 	producer, err := golang.NewProducer(&golang.Config{
// 		Endpoint: endpoint,
// 		Credentials: &credentials.SessionCredentials{
// 			AccessKey:    "",
// 			AccessSecret: "",
// 		},
// 	}, golang.WithTopics(topic))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return producer, nil
// }

func NewProducer(config *golang.Config, topic string) (golang.Producer, error) {
	producer, err := golang.NewProducer(config, golang.WithTopics(topic))
	if err != nil {
		return nil, err
	}
	return producer, nil
}
