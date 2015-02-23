# Go-bindings for C editline library

[Editline][] is a line editing library for UNIX.  It is call compatible with GNU readline, but is a fraction of the size and has fewer features.

[editline]: https://github.com/troglobit/editline

## Usage

        $ go get github.com/pm5/go-editline

You also need C editline library installed.  Then

```Go
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
```

## License

MIT License: <http://pm5.mit-license.org/>
