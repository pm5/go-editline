# Go-bindings for C editline library

## Usage

```golang
package main

import (
	"github.com/pm5/go-editline"
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
```

## License

MIT License: <http://pm5.mit-license.org/>
