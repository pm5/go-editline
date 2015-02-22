package editline

/*
#cgo LDFLAGS: -ledit -lc
#include <editline/readline.h>
#include <stdlib.h>
#include <string.h>
*/
import "C"
import "unsafe"
import (
	"fmt"
	"io"
)

func ReadLine(p string) (string, error) {
	cp := C.CString(p)
	defer C.free(unsafe.Pointer(cp))
	line := C.readline(cp)
	if line == nil {
		return "", io.EOF
	}
	return C.GoString(line), nil
}

func AddHistory(l string) {
	cl := C.CString(l)
	defer C.free(unsafe.Pointer(cl))
	C.add_history(cl)
}

func ReadHistory(filename string) error {
	cfn := C.CString(filename)
	defer C.free(unsafe.Pointer(cfn))
	r := C.read_history(cfn)
	if r != 0 {
		return fmt.Errorf("Cannot read history file `%s`: %s\n", filename, C.GoString(C.strerror(r)))
	}
	return nil
}
