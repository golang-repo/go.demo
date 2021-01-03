/**
 * @Author: v_bivwei
 * @Description:
 * @Date: 2021/1/1 20:11
 */
package main

import (
    "fmt"
    "runtime"
    "strconv"
    "strings"
)

func main() {
    fmt.Println(GoID())
}

func GoID() int {
    var buf [64]byte
    n := runtime.Stack(buf[:], false)
    fmt.Println("nnn ", string(buf[:n]))
    // 得到id字符串
    idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
    id, err := strconv.Atoi(idField)
    if err != nil {
        panic(fmt.Sprintf("cannot get goroutine id: %v", err))
    }
    return id
}
