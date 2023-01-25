package utl

import (
	"fmt"
	"log"
	"runtime"
)

func Log(err error) {
	pc, file, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(pc).Name()
	log.Println(fmt.Sprintf("***ERROR***\nFile  :\t%s\nFunc  :\t%s\nLine  :\t%d\nError :\n%s", file, name, line, err.Error()))
}
