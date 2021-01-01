/**
 * @Author: v_bivwei
 * @Description:
 * @Date: 2021/1/1 13:11
 */
package main

import "fmt"

func main() {
    var (
        list = []int{1, 52, 14, 3, 8}
    )
    ret := quickSort(list)
    fmt.Println("快排结果为: ", ret)
}

func quickSort(list []int) []int {
    if len(list) < 2 {
        return list
    }

    var (
        pivot       = list[0]
        lessList    = make([]int, 0)
        greaterList = make([]int, 0)
    )

    for _, v := range list {
        if v > pivot {
            greaterList = append(greaterList, v)
        } else if v == pivot {
            continue
        } else {
            lessList = append(lessList, v)
        }
    }

    return append(append(quickSort(lessList), pivot), quickSort(greaterList)...)

}
