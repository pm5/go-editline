# Go-bindings for C editline library

## Usage

        $ go get github.com/pm5/go-editline

Then

```Go
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
```

## License

MIT License: <http://pm5.mit-license.org/>
