package httputils

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {

	url := "https://webapp02.montcopa.org/eoc/cadinfo/livecad.asp"

	h := NewHTTP()

	r, err := h.Get(url)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}

	fmt.Println(string(r))

}
