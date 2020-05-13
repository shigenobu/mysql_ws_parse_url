/*
 MySQL UDF ws parse url
 */
package main

import (
	/*
		#cgo CFLAGS: -I/usr/include/mysql
		#include <stdio.h>
		#include <stdlib.h>
		#include <string.h>
		#include <mysql.h>
		#include <limits.h>
	*/
	"C"
	"unsafe"
	"./func"
)

// convert argc, argv into go structure
func argToGostrings(count C.uint, args **C.char, lengths *C.ulong) []string {
	// https://github.com/golang/go/wiki/cgo#turning-c-arrays-into-go-slices
	length := count
	argslice := (*[1 << 30]*C.char)(unsafe.Pointer(args))[:length:length]
	lengthsslice := (*[1]C.ulong)(unsafe.Pointer(lengths))[:length:length]

	gostrings := make([]string, count)

	for i, s := range argslice {
		l := C.int(lengthsslice[i])
		gostrings[i] = C.GoStringN(s, l)
	}
	return gostrings
}

//export ws_parse_url_init
func ws_parse_url_init(initid *C.UDF_INIT, args *C.UDF_ARGS, message *C.char) C.my_bool {
	length := int(args.arg_count)
	if length != 1 {
		return 1
	}

	typeslice := (*[1]C.uint)(unsafe.Pointer(args.arg_type))[:length:length]
	if typeslice[0] != C.STRING_RESULT {
		return 1
	}

	return 0
}

//export ws_parse_url
func ws_parse_url(initid *C.UDF_INIT, args *C.UDF_ARGS, result *C.char, length *C.ulong, isNull *C.char, error *C.char) *C.char {
	C.free(unsafe.Pointer(initid.ptr))

	argsString := argToGostrings(args.arg_count, args.args, args.lengths)

	// -----
	// ここが処理本体
	rslt := _func.ParseUrl(argsString[0])
	// -----

	initid.ptr = C.CString(rslt)
	*length = C.ulong(len(rslt))
	return initid.ptr
}

//export ws_parse_url_deinit
func ws_parse_url_deinit(initid *C.UDF_INIT) {
	C.free(unsafe.Pointer(initid.ptr))
}

func main() {
}
