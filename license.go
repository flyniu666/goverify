package goverify

/*
#cgo CFLAGS: -I .
#cgo CFLAGS: -I /usr/lib/java/include
#cgo CFLAGS: -I /usr/lib/java/linux/include
#cgo LDFLAGS: -L /var/lib -ljni
#cgo LDFLAGS: -L /usr/lib/java/jre/lib/amd64/server  -ljvm
#include <jni.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include "jni.h"
*/
import "C"

import "fmt"

func verify(srvNum int, cpuNum int memNum int) int {
//    Test("./verifyparam.properties")
  cstr := C.CString("/var/lib/verifyparam.properties")
  ret := C.verify(cstr, srvNum, cpuNum, memNum)
  fmt.Println("ret is   ", ret)
  return ret
}
