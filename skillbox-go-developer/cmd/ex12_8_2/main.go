package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	path = "../ex12_8_1/log.txt"
)

func main() {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}

	finfo, err := f.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("open file - %v, with size - %v\n", path, finfo.Size())

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		txt := sc.Text()
		fmt.Println(txt)
	}
}
