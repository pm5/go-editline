package main

import (
	".."
	"fmt"
)

func main() {
	var line string
	for {
		line = editline.ReadLine("> ")
		fmt.Println(line)
		editline.AddHistory(line)
	}
}
