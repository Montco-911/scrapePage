package pubSub

import "testing"

func TestProd(t *testing.T) {
	msg := NewMSG()
	msg.SetAddress("a")
	msg.Prod([]byte("test"))
}
