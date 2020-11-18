package pubSub

import (
	"context"
	"fmt"
	kafka "github.com/segmentio/kafka-go"
	"log"
	"time"
)

type MSG struct {
	address             string
	topic               string
	partition int
	errorMsgCount       int
	donTDisplayAfterNum int

}

func NewMSG() * MSG {
	m := &MSG{}
	m.topic = "activeAlerts"
	m.partition = 0
	m.address = "localhost:29092"
	m.donTDisplayAfterNum = 3
	return m
}

func (m *MSG)SetAddress(address string) {
    m.address = address
}
func (m *MSG)SetTopic(topic string) {
	m.topic = topic
}



func (m *MSG)Prod(msg []byte) {

	conn, err := kafka.DialLeader(context.Background(), "tcp", m.address, m.topic, m.partition)
	if err != nil {
		m.errorMsgCount+=1
		if m.errorMsgCount < m.donTDisplayAfterNum {
			log.Printf("failed to dial leader:%v\n", err)
		}
		return
	}

	conn.SetWriteDeadline(time.Now().Add(6*time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: msg},
	)
	if err != nil {
		m.errorMsgCount+=1
		if m.errorMsgCount < m.donTDisplayAfterNum {
			log.Printf("failed to write messages:%v\n", err)
		}

	}

	if err := conn.Close(); err != nil {
		m.errorMsgCount+=1
		if m.errorMsgCount < m.donTDisplayAfterNum {
			log.Printf("failed to close writer:%v\n", err)
		}
	}
}

func (msg *MSG)Reader(offset int64) {
	// make a new reader that consumes from topic-A, partition 0, at offset 42
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{msg.address},
		Topic:     msg.topic,
		Partition: 0,
		MinBytes:  1e3, // 1KB
		MaxBytes:  10e6, // 10MB
	})
	r.SetOffset(offset)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		log.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Printf("failed to close reader:%s\n", err)
	}
}

func (m *MSG)CreateTopic() {
	_, err := kafka.DialLeader(context.Background(), "tcp", m.address, m.topic, 0)
	if err != nil {
		fmt.Printf("err: %v\n",err)
	}
}