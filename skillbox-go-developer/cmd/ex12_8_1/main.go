package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	var (
		count int
	)

	f, err := os.Create("log.txt")
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("~: ")
		sc.Scan()
		str := sc.Text()

		if str == "exit" {
			break
		}

		count++
		date := time.Now().Format("2006-01-02 15:04:05")
		f.WriteString(fmt.Sprintf("%v %v %v\n", count, date, str))
	}
}
