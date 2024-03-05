package mq

import (
	"os"
	"time"

	"github.com/apache/rocketmq-clients/golang/v5"
)

func init() {
	os.Setenv("mq.consoleAppender.enabled", "true")
	golang.ResetLogger()
}

var (
	awaitDuration = time.Second * 5
)

func NewConsumer(endpoint string, topic string, consumerGroup string) (golang.Consumer, error) {
	consumer, err := golang.NewSimpleConsumer(&golang.Config{
		Endpoint:      endpoint,
		ConsumerGroup: consumerGroup,
	},
		golang.WithAwaitDuration(awaitDuration),
		golang.WithSubscriptionExpressions(map[string]*golang.FilterExpression{
			topic: golang.SUB_ALL,
		}),
	)

	if err != nil {
		return nil, err
	}

	return consumer, nil
}
