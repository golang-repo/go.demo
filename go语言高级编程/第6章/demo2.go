/**
 * @Author: v_bivwei
 * @Description:
 * @Date: 2021/1/2 18:56
 */
package main

import (
    "fmt"
    "unsafe"
)

type Programmer1 struct {
    name     string
    age      int
    language string
}

func main() {
    p := Programmer1{
        "stefno",
        18,
        "go",
    }
    fmt.Println(p)
    lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.language)))
    *lang = "Golang1111"
    fmt.Println(p)
}
