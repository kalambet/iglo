package iglo

/*
#cgo CPPFLAGS: -I${SRCDIR}/drafter/src -I${SRCDIR}/drafter/ext/snowcrash/src -I${SRCDIR}/drafter/ext/snowcrash/ext/markdown-parser/src -I${SRCDIR}/drafter/ext/snowcrash/ext/markdown-parser/ext/sundown/src -I${SRCDIR}/drafter/ext/sos/src -I${SRCDIR}/drafter/ext/cmdline
#cgo LDFLAGS: -L${SRCDIR}/drafter/build/out/Release -lsos -lsundown -lsnowcrash -lmarkdownparser -ldrafter
#include <stdlib.h>
#include "snowcrash.h"
#include "cdrafter.h"
*/
import "C"

import (
	"encoding/json"
	"io"
	"io/ioutil"
)
import "unsafe"

// ParseJSON ...
func ParseJSON(r io.Reader) (*API, error) {
	api := new(API)
	err := json.NewDecoder(r).Decode(&api)

	if err != nil {
		return nil, err
	}

	return api, nil
}

// ParseMarkdown ...
func ParseMarkdown(r io.Reader) ([]byte, error) {

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	length := len(b)
	source := C.CString(string(b[:length]))
	var result *string

	code := int(C.drafter_c_parse(
		source,
		C.int(0),
		C.int(0),
		(**C.char)(unsafe.Pointer(result)),
	))

	fmt.Ptintf("%#v", result)

	defer C.free(unsafe.Pointer(result))

	return nil, code
}
