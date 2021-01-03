/**
 * @Author: v_bivwei
 * @Description:
 * @Date: 2021/1/2 18:01
 */
package main

import (
    "fmt"
    "unsafe"
)

type Programmer struct {
    name     string
    language string
}

// https://www.sohu.com/a/319106990_657921
func main() {
    p := Programmer{"stefno", "go"}
    fmt.Println(p)
    name := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.name)))
    *name = "qcrao1111111"
    lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.language)))
    *lang = "Golang"
    fmt.Println(p)
}
