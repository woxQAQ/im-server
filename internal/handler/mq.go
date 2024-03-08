package handler

import (
	"context"
	"github.com/apache/rocketmq-clients/golang/v5"
	"github.com/panjf2000/ants/v2"
	"go.uber.org/zap"
)

type Connector interface {
	Receive(ctx context.Context) error
	Producer(ctx context.Context, message *golang.Message) error
}

type mqConnector struct {
	// MqConsumer is rockerMq consumer
	// consumer is used to get message from gateway
	MqConsumerList []golang.SimpleConsumer

	// GatewayProducer is rockerMq producer
	// producer is used to write message to database
	GatewayProducer golang.Producer

	goroutinePool *ants.Pool

	messageChan chan []byte
}

var _ Connector = (*mqConnector)(nil)

func (c *mqConnector) Producer(ctx context.Context, message *golang.Message) error {
	//defer func() {
	//}()

	err := c.GatewayProducer.Start()
	if err != nil {
		return err
	}

	defer c.GatewayProducer.GracefulStop()

	c.GatewayProducer.SendAsync(ctx, message, func(ctx context.Context, receipts []*golang.SendReceipt, err error) {
		if err != nil {
			zap.S().Info(err)
		}
	})

	return nil
}

func (c *mqConnector) Receive(ctx context.Context) error {
	//err := c.GatewayProducer.Start()
	//if err != nil {
	//	zap.S().Fatal(err)
	//}

	defer func() {
		//_ = c.GatewayProducer.GracefulStop()
		close(c.messageChan)
	}()

	for _, consumer := range c.MqConsumerList {
		err := c.goroutinePool.Submit(func() {
			c.consumerRun(ctx, consumer)
		})
		if err != nil {
			return nil
		}
	}
	return nil
}

//func (c *mqConnector) producerRun(producer golang.Producer, wg *sync.WaitGroup) {
//	//todo
//}

func (c *mqConnector) consumerRun(ctx context.Context, consumer golang.SimpleConsumer) {
	err := consumer.Start()
	if err != nil {
		zap.S().Fatal(err)
	}
	defer func() {
		consumer.GracefulStop()
	}()
	for {
		data, err := consumer.Receive(context.TODO(), maxMessageNum, invisibleDuration)
		if err != nil {
			zap.S().Infoln(err)
		}

		for _, d := range data {
			c.messageChan <- d.GetBody()
			err := consumer.Ack(context.TODO(), d)
			if err != nil {
				zap.S().Infoln(err)
			}
		}
	}

}
