package activeIncident

import (
	"strings"
	"testing"
)

func TestActiveJSON(t *testing.T) {
	result := ActiveJSON()
	if !strings.Contains(string(result),"9252418") {
		t.Fatalf("fail.. expected eid ")

	}
}