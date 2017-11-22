package bsdp

/*
#cgo CFLAGS: -I../c/bsdiff -I../c/bspatch -I../c/bzip2
#include "bsdiff.c"
#include "bspatch.c"
#include "bzlib.c"
#include "compress.c"
#include "crctable.c"
#include "randtable.c"
#include "blocksort.c"
#include "huffman.c"
#include "decompress.c"
*/
import "C"
import "unsafe"
import "errors"

func Diff(oldfile, newfile, patchfile string) error {
    coldfile := C.CString(oldfile)
    cnewfile := C.CString(newfile)
    cpatchfile := C.CString(patchfile)
    
    var errBytes [10245]byte
    cerr := (*C.char)(unsafe.Pointer(&errBytes[0]))

    defer C.free(unsafe.Pointer(coldfile))
    defer C.free(unsafe.Pointer(cnewfile))
    defer C.free(unsafe.Pointer(cpatchfile))

    C.bsdiff(cerr, coldfile, cnewfile, cpatchfile)
    
    err := C.GoString(cerr)
    if err != "" {
        return errors.New(err)
    }
    return nil
}

func Patch(oldfile, newfile, patchfile string) error {
    coldfile := C.CString(oldfile)
    cnewfile := C.CString(newfile)
    cpatchfile := C.CString(patchfile)
    
    var errBytes [10245]byte
    cerr := (*C.char)(unsafe.Pointer(&errBytes[0]))

    defer C.free(unsafe.Pointer(coldfile))
    defer C.free(unsafe.Pointer(cnewfile))
    defer C.free(unsafe.Pointer(cpatchfile))

    C.bspatch(cerr, coldfile, cnewfile, cpatchfile)
    
    err := C.GoString(cerr)
    if err != "" {
        return errors.New(err)
    }
    return nil
}