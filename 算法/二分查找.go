/**
 * @Author: v_bivwei
 * @Description:
 * @Date: 2020/12/30 23:02
 */
package main

import (
    "fmt"
)

var (
    NO_FOUND_INDEX2 = -1
)

func main() {
    list := []int{1, 2, 3, 7, 8, 58, 98}
    target := 98
    index := binarySearch(list, target)
    if index == NO_FOUND_INDEX2 {
        fmt.Printf("数组 %+v 中没有找到没有找到 %d \n", list, target)
    } else {
        fmt.Printf("index = %d ,利用索引拿到的值为 %d ,实际的值为 %d ", index, list[index], target)
    }
}

func binarySearch(list []int, target int) (index int) {
    low := 0
    high := len(list) - 1

    // 只要范围没有缩小到只包含一个元素
    // 当只剩下一个元素时, low = high = 0,此时 mid 也等于 0,list[0] 也就是数组中的那一个元素
    for low <= high {
        mid := (low + high) / 2 // go 会自动向下取整
        if list[mid] == target {
            return mid
        } else if list[mid] > target {
            high = mid - 1
        } else if list[mid] < target {
            low = mid + 1
        } else {
            break
        }
    }

    return NO_FOUND_INDEX2

}
