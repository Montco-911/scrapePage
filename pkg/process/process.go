package process

import (
	"fmt"
	"github.com/Montco-911/scrapePage/pkg/httputils"
	"github.com/Montco-911/scrapePage/pkg/parse"
	"log"
	"os"
)

func WriteJson(filename string) error {

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		return err
	}
	defer f.Close()

	a, err := GetJson()
	if err != nil {
		log.Printf("Error in json")
	}
	if _, err = f.WriteString(string(a)); err != nil {
		return err
	}
	return nil
}

//TODO: Fix me
func GetJson() ([]byte, error) {
	c, a, err := BuildDb()
	if err != nil {
		log.Fatalf("No build")
	}

	return parse.ToJson(c, a)

}
func ShowJson() {
	a, err := GetJson()
	if err != nil {
		log.Printf("Error in json")
	}
	println(string(a))
}

func BuildDb() ([]map[string]string, [][]string, error) {

	callTable := []map[string]string{}
	arriveTable := [][]string{}

	url := "https://webapp02.montcopa.org/eoc/cadinfo/livecad.asp"
	h := httputils.NewHTTP()
	r, err := h.Get(url)

	if err != nil {
		return nil, nil, err
	}

	result, link, err := parse.Tag(string(r))
	if err != nil {
		return nil, nil, err
	}

	for _, result := range result {
		callTable = append(callTable, parse.Strip(result))
	}

	for _, l := range link {
		r, err = h.Get(parse.GetDetail(l))
		if err != nil {
			return callTable, nil, err
		}

		arrive, err := parse.GetTable(string(r))
		if err != nil {
			return callTable, nil, err
		}
		arriveTable = append(arriveTable, arrive)

	}

	return callTable, arriveTable, err

}

func Show() {
	c, a, err := BuildDb()
	if err != nil {
		log.Fatalf("No build")
	}
	for i, m := range c {
		for k, v := range m {
			fmt.Printf("%v: %v\n", k, v)
		}
		fmt.Printf("Status: %v\n\n", a[i])
	}
}
