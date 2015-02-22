package main

import (
	".."
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		for _, fn := range os.Args[1:] {
			err := editline.ReadHistory(fn)
			if err != nil {
				fmt.Print(err)
				os.Exit(-1)
			}
		}
	}
	for {
		line, err := editline.ReadLine("> ")
		if err == io.EOF {
			break
		}
		fmt.Println(line)
		editline.AddHistory(line)
	}
}
