package parse

import (
	"encoding/json"
	"golang.org/x/net/html"
	"log"
	"strings"
	"time"
)

type DB struct {
	r map[string]string
	v []string
}

func Tag(s string) ([]string, []string, error) {
	doc, err := html.Parse(strings.NewReader(s))
	r := []string{}
	l := []string{}
	if err != nil {
		return r, l, err
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {

					if strings.Contains(a.Val, "map.asp?type=") {
						// fmt.Println(a.Val)
						r = append(r, a.Val)
					} else if strings.Contains(a.Val, "livecad") {
						l = append(l, a.Val)
					}

					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return r, l, err
}

func Strip(s string) map[string]string {

	//fmt.Printf("%v\n", s)
	m := map[string]string{}
	s = CleanUp(s)
	for _, v := range strings.Split(s, "&") {
		ss := strings.Split(v, "=")
		if len(ss) == 2 {
			//fmt.Printf("M: %s, %s\n", ss[0], ss[1])
			m[ss[0]] = ss[1]
		}

	}
	return m
}

func CleanUp(s string) string {
	s = strings.Replace(s, "livecadcomments-fireems.asp?eid", "eid", -1)
	s = strings.Replace(s, "map.asp?type", "type", -1)
	s = strings.Replace(s, "<br>", " ", -1)
	s = strings.Replace(s, " @ ", " ", -1)
	return s
}

func GetDetail(purl string) string {
	url := "https://webapp02.montcopa.org/eoc/cadinfo/" + purl
	return strings.Replace(url, " ", "%20", -1)
}

func GetTable(s string) ([]string, error) {
	doc, err := html.Parse(strings.NewReader(s))
	r := []string{}

	if err != nil {
		return r, err
	}
	var f func(*html.Node)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "table" {

		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {

			if c.Data == "td" {

				if c.FirstChild.Data == "b" {
					//c = c.FirstChild
					return
				}

				if c.FirstChild.Data == "font" {
					r = append(r, c.FirstChild.FirstChild.Data)
				} else {
					r = append(r, c.FirstChild.Data)
				}

			}

			f(c)
		}
	}
	f(doc)

	return r, nil
}

func GetTableV2(s string) ([]string, error) {
	doc, err := html.Parse(strings.NewReader(s))
	r := []string{}

	if err != nil {
		return r, err
	}
	var f func(*html.Node)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "table" {

		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {

			if c.Data == "td" {

				if c.FirstChild.Data == "b" {
					//c = c.FirstChild
					return
				}

				if c.FirstChild.Data == "font" {
					r = append(r, c.FirstChild.FirstChild.Data)
				} else {
					r = append(r, c.FirstChild.Data)
				}

			}

			f(c)
		}
	}
	f(doc)

	return r, nil
}

func ToJson(call []map[string]string, status [][]string) ([]byte, error) {
	type Calls struct {
		Call   map[string]string
		Status []string
	}

	calls := []*Calls{}

	if len(status) < len(call) {
		log.Printf("len(status) < len(call)\n")
		for i := len(status); i < len(call); i++ {
			status = append(status, []string{})
		}
	}

	for i, v := range call {
		nt := new(Calls)
		nt.Call = v
		nt.Status = status[i]
		calls = append(calls, nt)
	}

	type DB struct {
		Calls     []*Calls
		TimeStamp time.Time
	}

	return json.Marshal(DB{calls, time.Now()})

}
