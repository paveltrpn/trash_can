package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func listDir(path string) (out []string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		out = append(out, path+"/"+file.Name())
	}

	return out
}

func main() {
	dir := listDir(os.Args[1])

	for _, fname := range dir {
		old_fname := fname

		id := strings.Index(fname, "(zaycev.net)")
		if id != -1 {
			fname = strings.Replace(fname, "_(zaycev.net)", "", id)
		}

		for i := 0; i < len(fname); i++ {
			if fname[i] == '_' {
				fname = strings.Replace(fname, "_", " ", i)
			}
		}

		os.Rename(old_fname, fname)

		fmt.Println(fname)
	}
}
