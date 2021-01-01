/**
 * @Author: v_bivwei
 * @Description:
 * @Date: 2020/12/31 17:22
 */
package main

import "fmt"

func main() {
    diGui(10)
}

func diGui(i int) {
    fmt.Println("i = ", i)
    if i < 1 {
        return
    }
    diGui(i - 1)
}
