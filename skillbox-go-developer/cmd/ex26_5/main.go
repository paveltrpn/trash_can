package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	// if cat launch with zero arguments then echo user enter
	// in for loop
	if len(os.Args) < 2 {
		for {
			var foo string
			fmt.Scan(&foo)
			fmt.Println(foo)
		}
	}

	var (
		files    []*os.File
		notExist bool = false
	)

	for i := 1; i < len(os.Args); i++ {
		file, err := os.OpenFile(os.Args[i], os.O_RDWR|os.O_APPEND, 0x755)
		if err != nil {
			fmt.Printf("no such file - %v\n", os.Args[i])
			// set flag if at least one file not exist
			notExist = true
		}
		files = append(files, file)
	}

	// terminate if at least one file not exist
	if notExist {
		os.Exit(0)
	}

	switch len(files) {
	case 1:
		buf := new(bytes.Buffer)
		buf.ReadFrom(files[0])
		fmt.Printf("first file:\n%v\n", buf.String())
	case 2:
		bufFirst := new(bytes.Buffer)
		bufFirst.ReadFrom(files[0])

		bufSecond := new(bytes.Buffer)
		bufSecond.ReadFrom(files[1])

		var str = []string{bufFirst.String(), bufSecond.String()}
		result := strings.Join(str, "")
		fmt.Printf("%v", result)
	case 3:
		bufFirst := new(bytes.Buffer)
		bufFirst.ReadFrom(files[0])

		bufSecond := new(bytes.Buffer)
		bufSecond.ReadFrom(files[1])

		var str = []string{bufFirst.String(), bufSecond.String()}
		result := strings.Join(str, "")
		files[2].WriteString(result)
	}
}
