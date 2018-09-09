package main


import "C"

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("./piumago.profile")
    if err != nil {
        fmt.Println(err)
    }
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	OptimizeFromDirWrapper(C.CString("../images"), 100, 50);
}
