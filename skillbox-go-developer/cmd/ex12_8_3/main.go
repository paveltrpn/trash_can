package main

import (
	"log"
	"os"
)

const (
	fname = "tmp.txt"
)

func main() {
	f, err := os.Create(fname)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	f.Close()

	if err := os.Chmod(fname, 0444); err != nil {
		log.Fatal(err)
	}

	f, err = os.Open(fname)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	if _, err := f.WriteString("foo bar"); err != nil {
		log.Fatalln(err)
	}
}
