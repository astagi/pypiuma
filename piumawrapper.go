package main

/*
struct PiumaResult {
	char* path;
	char* message;
} ;
*/
import "C"

import (
	_"fmt"
	"unsafe"
	"io/ioutil"
	"path/filepath"
)


//export OptimizeListWrapper
func OptimizeListWrapper(path **C.char, pathElements int, width uint, height uint) (C.struct_PiumaResult) {
	cStrings := (*[1 << 30]*C.char)(unsafe.Pointer(path))[:pathElements:pathElements]
	for _, path := range cStrings {
		Optimize(C.GoString(path), width, height)
	}
	return C.struct_PiumaResult{path: nil, message: nil}
}

//export OptimizeFromDirWrapper
func OptimizeFromDirWrapper(path *C.char, width uint, height uint) (C.struct_PiumaResult) {
	var targetPath = C.GoString(path)
	files, _ := ioutil.ReadDir(targetPath)
	for _, f := range files {
		Optimize(filepath.Join(targetPath, f.Name()), width, height)
	}
	return C.struct_PiumaResult{path: nil, message: nil}
}

//export OptimizeWrapper
func OptimizeWrapper(path *C.char, width uint, height uint) (C.struct_PiumaResult) {
	destPath, err := Optimize(C.GoString(path), width, height)
	if err == nil{
		return C.struct_PiumaResult{path: C.CString(destPath), message: nil}
	} else {
		return C.struct_PiumaResult{path: nil, message: C.CString(err.Error())}
	}
}

func main() {
}
