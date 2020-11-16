package activeIncident

import (
	"encoding/json"
	"github.com/Montco-911/scrapePage/test-fixtures/activeIncident"
	"testing"
)

func Test_Json(t *testing.T) {
	b := activeIncident.ActiveJSON()
	var a ActiveIncident
	json.Unmarshal(b, &a)
	if len(a.Calls) != 15 {
		t.Fatalf("Can't Unmarshal json")
	}

}