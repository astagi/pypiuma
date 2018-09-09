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
	"sync"
	"unsafe"
	"github.com/karrick/godirwalk"
)


//export OptimizeListWrapper
func OptimizeListWrapper(path **C.char, pathElements int, width uint, height uint) (C.struct_PiumaResult) {
	var wg sync.WaitGroup
	cStrings := (*[1<<30]*C.char)(unsafe.Pointer(path))[:pathElements:pathElements]
	for _, path := range cStrings {
		wg.Add(1)
		go func(filePath string, width uint, height uint) {
			defer wg.Done()
			Optimize(filePath, width, height)
		}(C.GoString(path), width, height)
	}
	wg.Wait()
	return C.struct_PiumaResult{path: nil, message: nil}
}

//export OptimizeFromDirWrapper
func OptimizeFromDirWrapper(path *C.char, width uint, height uint) (C.struct_PiumaResult) {
	var wg sync.WaitGroup
	godirwalk.Walk(C.GoString(path), &godirwalk.Options {
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			wg.Add(1)
			go func(filePath string, width uint, height uint) {
				defer wg.Done()
				Optimize(filePath, width, height)
			}(osPathname, width, height)
			return nil
		},
		Unsorted: true,
	})

	wg.Wait()
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
