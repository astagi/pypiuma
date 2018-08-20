package main

/*
struct PiumaResult {
  char* path;
  char* message;
} ;
*/
import "C"


//export OptimizeWrapper
func OptimizeWrapper(path *C.char, width uint, height uint) (C.struct_PiumaResult) {
	destPath, err := Optimize(C.GoString(path), width, height)
	if err == nil{
		return C.struct_PiumaResult{path: C.CString(destPath), message: nil}
	} else {
		return C.struct_PiumaResult{path: nil, message: C.CString("Orror")}
	}
}

func main() {
}
