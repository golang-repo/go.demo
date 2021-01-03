/**
 * @Author: v_bivwei
 * @Description:
 * @Date: 2021/1/1 19:35
 */
package main

import (
    "fmt"
    "sync"
)

func main() {
    var (
        mu    sync.Mutex
        count = 0
    )
    // 使用WaitGroup等待10个goroutine完成
    var wg sync.WaitGroup
    wg.Add(10)
    for i := 0; i < 10; i++ {
        go func() {
            defer wg.Done()
            // 对变量count执行10次加1
            for j := 0; j < 100000; j++ {
                mu.Lock()
                count++
                mu.Unlock()
            }
        }()
    }
    // 等待10个goroutine完成
    wg.Wait()
    fmt.Println(count)
}
