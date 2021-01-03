/**
 * @Author: v_bivwei
 * @Description:
 * @Date: 2021/1/1 19:56
 */
package main

import (
    "fmt"
    "sync"
)

type Counter1 struct {
    Count int64
    mu    sync.Mutex
}

func main() {
    var (
        count Counter1
    )
    // 使用WaitGroup等待10个goroutine完成
    var wg sync.WaitGroup
    wg.Add(10)
    for i := 0; i < 10; i++ {
        go func(count *Counter1) {
            defer wg.Done()
            // 对变量count执行10次加1
            for j := 0; j < 100000; j++ {
                count.mu.Lock()
                count.Count++
                count.mu.Unlock()
            }
        }(&count)
    }
    // 等待10个goroutine完成
    wg.Wait()
    fmt.Println(count.Count)
}
