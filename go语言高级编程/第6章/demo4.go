/**
 * @Author: v_bivwei
 * @Description:
 * @Date: 2021/1/2 22:25
 */
package main

import (
    "fmt"
    "unsafe"
)

func main() {
    s := make([]int, 9, 20)
    var Len = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Sizeof(int(0))))
    fmt.Println(Len, len(s)) // 9 9
}
