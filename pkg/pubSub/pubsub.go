package pubSub

import (
	"context"
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
