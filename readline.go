package editline

/*
#cgo LDFLAGS: -ledit
#include <editline/readline.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"
import "io"

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
