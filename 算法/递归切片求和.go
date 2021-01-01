/**
 * @Author: v_bivwei
 * @Description:
 * @Date: 2020/12/31 18:01
 */
package main

import (
    "fmt"
)

func main() {
    var (
        arr = []int{1, 2, 3, 4}
        sum = 0
    )

    ret := diGuiSum(arr, sum)
    fmt.Println("结果为: ", ret)

}

func diGuiSum(arr []int, sum int) int {
    if len(arr) == 0 {
        return sum
    }

    return diGuiSum(arr[1:], sum+arr[0])
}
