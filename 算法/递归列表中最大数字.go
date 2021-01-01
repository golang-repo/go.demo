/**
 * @Author: v_bivwei
 * @Description:
 * @Date: 2021/1/1 0:11
 */
package main

import "fmt"

func main() {
    var (
        list = []int{1, 3, 3, 111110, 14, 10000}
    )

    max := diGuiMaxInList(list, list[0])
    fmt.Println("最大值为 ", max)

}

func diGuiMaxInList(list []int, max int) int {
    if len(list) == 0 {
        return max
    }

    if list[0] > max {
        max = list[0]
        return diGuiMaxInList(list[1:], max)
    } else {
        return diGuiMaxInList(list[1:], max)
    }

}
