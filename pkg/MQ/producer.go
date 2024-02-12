package MQ

import "github.com/segmentio/kafka-go"

type Producer struct {
	addr   []string
	topic  string
	writer *kafka.Writer
}

func NewKafkaProducer(addr []string, topic string) (*Producer, error) {
	p := Producer{
		addr:  addr,
		topic: topic,
	}

	// if the env variable is not set, use the default value
	p.addr = readAddrFromEnv(p.addr)
	writer := kafka.Writer{
		Addr:         kafka.TCP(p.addr...),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
	}
	p.writer = &writer
	return &p, nil
}
