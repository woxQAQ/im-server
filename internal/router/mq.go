package router

import (
	"context"
	"sync"

	"github.com/apache/rocketmq-clients/golang/v5"
	"github.com/panjf2000/ants/v2"
	"go.uber.org/zap"
)

type Connector interface {
	Run()
}

type mqConnector struct {
	// MqConsumer is rockerMq consumer
	// consumer is used to get message from gateway
	MqConsumerList []golang.SimpleConsumer

	// MqProducer is rockerMq producer
	// producer is used to write message to database
	MqProducer golang.Producer

	goroutinePool *ants.Pool

	messageChan chan *golang.MessageView
}

var _ Connector = (*mqConnector)(nil)

func (c *mqConnector) Run() {
	var wg *sync.WaitGroup
	err := c.MqProducer.Start()
	if err != nil {
		zap.S().Fatal(err)
	}

	defer c.MqProducer.GracefulStop()

	for i := 0; i < consumerNum; i++ {
		wg.Add(1)
		err := c.goroutinePool.Submit(func() {
			c.consumerRun(c.MqConsumerList[i], wg)
		})
		if err != nil {
			zap.S().Fatal(err)
		}
	}

	wg.Wait()
}

func (c *mqConnector) producerRun(producer golang.Producer, wg *sync.WaitGroup) {
	//todo
}

func (c *mqConnector) consumerRun(consumer golang.SimpleConsumer, wg *sync.WaitGroup) {
	err := consumer.Start()
	if err != nil {
		zap.S().Fatal(err)
	}
	defer func() {
		consumer.GracefulStop()
		wg.Done()
	}()
	for {
		data, err := consumer.Receive(context.TODO(), maxMessagenum, invisidibleDuration)
		if err != nil {
			zap.S().Info(err)
		}
		for _, d := range data {
			c.messageChan <- d
			consumer.Ack(context.TODO(), d)
		}
	}

}
