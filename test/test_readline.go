package main

import (
	"fmt"
	"github.com/pm5/go-editline"
)

func main() {
	var line string
	for {
		line = editline.ReadLine("> ")
		fmt.Println(line)
		editline.AddHistory(line)
	}
}
