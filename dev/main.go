package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func WriteAppend(file, txt string) {
	f, err := os.OpenFile(file,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("WriteAppend: %v\n", err)
	}
	defer f.Close()
	f.Write([]byte(txt))
}

func Open() {
	dat, _ := ioutil.ReadFile("/etc/passwd")
	fmt.Print(string(dat))
}

func Read(file string, b []byte) (int, error) {
	f, err := os.OpenFile(file,
		os.O_RDWR, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return f.Read(b)
}

func main() {
	Open()
	WriteAppend("/data/junk", "stuff")
	b := make([]byte, 300)
	Read("/data/junk", b)
	fmt.Printf("b: %s\n", b)
}
