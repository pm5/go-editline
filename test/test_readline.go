package main

import (
	"fmt"
	"github.com/pm5/go-editline"
	"io"
)

func main() {
	for {
		line, err := editline.ReadLine("> ")
		if err == io.EOF {
			break
		}
		fmt.Println(line)
		editline.AddHistory(line)
	}
}
