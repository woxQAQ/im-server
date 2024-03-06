package mq

import (
	"context"
	"testing"

	"github.com/apache/rocketmq-clients/golang/v5"
)

func TestProducer(t *testing.T) {
	p, err := NewProducer("172.22.141.30:8081", "GATEWAY_WITH_ROUTE")
	if err != nil {
		t.Error(err)
	}

	err = p.Start()
	if err != nil {
		t.Error(err)
	}
	defer func() {
		p.GracefulStop()
	}()

	rmqMsg := golang.Message{
		Topic: "GATEWAY_WITH_ROUTE",
		Body:  []byte("helloworld"),
	}

	rmqMsg.SetTag("chatmsg")

	_, err = p.Send(context.TODO(), &rmqMsg)
	if err != nil {
		t.Error(err)
	}

}
