package parse

import (
	"fmt"
	test_fixtures "github.com/Montco-911/scrapePage/test-fixtures"
	"testing"
)

func TestGetTableV2(t *testing.T) {

	a, _ := GetTableV2(test_fixtures.Detail())
	for _, v := range a {
		fmt.Printf("%v\n", v)
	}
}
