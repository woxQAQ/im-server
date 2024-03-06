package router

import (
	"time"

	"github.com/apache/rocketmq-clients/golang/v5"
	"github.com/panjf2000/ants/v2"
	"github.com/woxQAQ/im-service/config"
	"github.com/woxQAQ/im-service/pkg/common/mq"
)

var (
	consumerNum               = 3
	maxMessagenum       int32 = 16
	invisidibleDuration       = time.Second * 32
)

type Router struct {
	// // GroupMsgControl is used to manage GroupMsg database
	// GroupMsgControl group.GroupMsgModel

	// // SingleMsgControl is used to manage SingleMsg database
	// SingleMsgControl single.SingleMessageModel

	// // GroupUserControl is used to manage GroupUserControl database
	// GroupUserControl group.GroupUsersModel

	// RpcHandler is used to router im message, and call rpc
	RpcHandler RpcRouterHandler

	*mqConnector
}

func NewRouter(config *config.RouterConfig) (*Router, error) {
	//todo

	var consumerList []golang.SimpleConsumer
	for i := 0; i < consumerNum; i++ {
		consumer, err := mq.NewConsumer(&config.RmqConfig, config.Topic.ConsumerTopic)
		if err != nil {
			return nil, err
		}
		consumerList = append(consumerList, consumer)
	}

	producer, err := mq.NewProducer(&config.RmqConfig, config.Topic.ProducerTopic)
	if err != nil {
		return nil, err
	}
	pool, err := ants.NewPool(1024, ants.WithNonblocking(true))
	if err != nil {
		return nil, err
	}
	return &Router{
		RpcHandler: newHandler(config.RpcConfig),
		mqConnector: &mqConnector{
			MqConsumerList: consumerList,
			MqProducer:     producer,
			goroutinePool:  pool,
		},
	}, nil
}
