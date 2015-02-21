package editline

/*
#cgo LDFLAGS: -ledit
#include <editline/readline.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

func ReadLine(p string) string {
	cp := C.CString(p)
	defer C.free(unsafe.Pointer(cp))
	return C.GoString(C.readline(cp))
}

func AddHistory(l string) {
	cl := C.CString(l)
	defer C.free(unsafe.Pointer(cl))
	C.add_history(cl)
}
