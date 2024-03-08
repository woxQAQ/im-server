package handler

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_message/msg"
	"github.com/woxQAQ/im-service/pkg/utils"
	"github.com/zeromicro/go-zero/zrpc"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"sync"
	"time"

	"github.com/apache/rocketmq-clients/golang/v5"
	"github.com/panjf2000/ants/v2"
	"github.com/woxQAQ/im-service/config"
	"github.com/woxQAQ/im-service/pkg/common/mq"
)

var (
	consumerNum             = 3
	maxMessageNum     int32 = 16
	invisibleDuration       = time.Second * 32
	retryTimes              = 3
)

type Handler struct {
	// // GroupMsgControl is used to manage GroupMsg database
	// GroupMsgControl group.GroupMsgModel

	// // SingleMsgControl is used to manage SingleMsg database
	// SingleMsgControl single.SingleMessageModel

	// // GroupUserControl is used to manage GroupUserControl database
	// GroupUserControl group.GroupUsersModel

	// RpcHandler is used to handler im message, and call rpc
	*mqConnector

	rmqProducerTopic string

	msg      msg.Msg
	validate *validator.Validate
	reqPool  sync.Pool
	Encoder
	utils.Retry
}

type RpcRouterHandler interface {
	SendMessage(context.Context, *SendRequest) ([]byte, error)
	DataAccept(context.Context, chan []byte)
}

func New(config *config.RouterConfig) (*Handler, error) {
	// init consumers
	var consumerList []golang.SimpleConsumer
	for i := 0; i < consumerNum; i++ {
		consumer, err := mq.NewConsumer(&config.RmqConfig, config.Topic.ConsumerTopic)
		if err != nil {
			return nil, err
		}
		consumerList = append(consumerList, consumer)
	}

	// init producers
	producer, err := mq.NewProducer(&config.RmqConfig, config.Topic.ProducerTopic)
	if err != nil {
		return nil, err
	}

	pool, err := ants.NewPool(1024, ants.WithNonblocking(true))
	if err != nil {
		return nil, err
	}
	handler := &Handler{
		mqConnector: &mqConnector{
			MqConsumerList:  consumerList,
			GatewayProducer: producer,
			goroutinePool:   pool,
		},
		rmqProducerTopic: config.Topic.ProducerTopic,
		validate:         validator.New(),
		msg:              msg.NewMsg(zrpc.MustNewClient(config.RpcConfig)),
		Encoder:          newGobEncoder(),
		reqPool: sync.Pool{New: func() any {
			return new(*SendRequest)
		}},
		Retry: utils.NewSimpleRetry(retryTimes, time.Second),
	}

	return handler, nil
}

func (h *Handler) Run() error {
	err := h.mqConnector.Receive(context.Background())
	if err != nil {
		return err
	}
	//var wg sync.WaitGroup
	//
	//wg.Done()
	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		h.DataAccept(ctx, h.messageChan)
	}()

	return nil
}

func (h *Handler) handlerMessage(context context.Context, data []byte) error {
	req := h.reqPool.Get().(*SendRequest)
	err := h.Decode(data, req)
	if err != nil {
		zap.S().Errorf("Decode error, err: %v", err)
	}

	for i := 0; i < retryTimes; i++ {
	}

	resp, err := h.SendMessage(context, req)
	if err != nil {
		return err
	}

	h.reqPool.Put(req)

	message := &golang.Message{
		Topic: h.rmqProducerTopic,
		Body:  resp,
	}

	err = h.Producer(context, message)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) DataAccept(context context.Context, messageChan chan []byte) {
	for {
		select {
		case data := <-messageChan:
			go func() {
				err := h.Do(func() error {
					err := h.handlerMessage(context, data)
					if err != nil {
						return err
					}
					return nil
				})

				if err != nil {
					zap.S().Errorln(err)
				}
			}()
		}
	}
}

func (h *Handler) SendMessage(ctx context.Context, data *SendRequest) ([]byte, error) {
	if err := h.validate.Struct(data); err != nil {
		return nil, err
	}

	req := msg.SendMessageReq{}
	getRPCReq(data, &req)

	resp, err := h.msg.SendMsg(ctx, &req)
	if err != nil {
		return nil, err
	}
	res, err := proto.Marshal(resp)
	if err != nil {
		return nil, err
	}

	return res, nil
}
