/**
 * @Author: v_bivwei
 * @Description:
 * @Date: 2020/12/30 21:36
 */
package main

import (
    "fmt"
    "time"
)

/**
main 函数是一个主协程
*/
func main() {
    fmt.Println("=====================")

    c1 := make(chan struct{}) // 无缓存的通道
    go func() {
        fmt.Println("c1....")
        time.Sleep(15 * time.Second)
        c1 <- struct{}{}
    }()

    fmt.Println("开始等 15s....")
    // 这里读取管道是阻塞的,因为子协程睡了 15s 后才往管道里写了一个无类型的匿名结构体
    // 所以这里的阻塞了 15s
    <-c1
    fmt.Println("15s 结束")

}
