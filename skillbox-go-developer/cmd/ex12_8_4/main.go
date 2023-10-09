package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	var (
		count  int
		buffer string
	)

	sc := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("~: ")
		sc.Scan()
		str := sc.Text()

		if str == "exit" {
			ioutil.WriteFile("log.txt", []byte(buffer), 0644)
			break
		}

		count++
		date := time.Now().Format("2006-01-02 15:04:05")
		buffer += fmt.Sprintf("%v %v %v\n", count, date, str)
	}

}
