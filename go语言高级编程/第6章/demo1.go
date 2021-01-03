/**
 * @Author: v_bivwei
 * @Description:
 * @Date: 2021/1/2 18:47
 */
package main

import (
    "fmt"
    "unsafe"
)

type X struct {
    state int32
    sema  uint32
}

func main() {
    x := X{
        state: 1,
        sema:  15,
    }

    dd := (*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Sizeof(uint32(0))))
    *dd = 25555
    fmt.Println(x)

}
