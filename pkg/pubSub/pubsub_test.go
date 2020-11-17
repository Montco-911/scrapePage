package pubSub

import "testing"

func TestProd(t *testing.T) {
	msg := NewMSG()
	msg.CreateTopic()
	//msg.SetAddress("a")
	msg.Prod([]byte("test"))
	//msg.Reader(0)
}
