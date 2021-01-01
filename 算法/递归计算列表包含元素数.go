/**
 * @Author: v_bivwei
 * @Description:
 * @Date: 2021/1/1 0:06
 */
package main

import "fmt"

func main() {
    var (
        list = []int{1, 3, 3, 4, 5, 5, 5, 5, 5, 5, 5}
        num  = 0
    )
    ret := diGuiListNum(list, num)
    fmt.Println("列表元素个数为: ", ret)
}

func diGuiListNum(list []int, num int) int {
    if len(list) == 0 {
        return num
    }

    return diGuiListNum(list[1:], num+1)

}
