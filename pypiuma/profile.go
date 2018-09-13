package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"

import (
    "fmt"
    "os"
    "runtime/pprof"
    "unsafe"
)

func main() {
    f, err := os.Create("./piumago.profile")
    if err != nil {
        fmt.Println(err)
    }
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()
    cs := C.CString("../images")
    defer C.free(unsafe.Pointer(cs))
    OptimizeFromDirWrapper(cs, 100, 50);
}
