package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

const (
	path = "../ex12_8_4/log.txt"
)

func main() {
	buffer, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(buffer))
}
