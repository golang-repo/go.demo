/**
 * @Author: v_bivwei
 * @Description:
 * @Date: 2021/1/1 16:30
 */
package main

import "fmt"

func main() {
    var (
        list = []int{2, 3, 7, 8, 10}
    )
    cfb(list)
}

func cfb(list []int) {
    for _, val := range list {
        for _, v := range list {
            fmt.Printf("%d * %d  ", val, v)
        }
        fmt.Println()
    }
}

// todo 用递归可以实现吗
func cfbV2(list []int) {

}
