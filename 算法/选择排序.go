/**
 * @Author: v_bivwei
 * @Description:
 * @Date: 2020/12/31 15:54
 */
package main

import "fmt"

func main() {
    var (
        arr    = []int{3, 4, 58, 1, 23, 11, 85, 47}
        newArr = make([]int, 0)
    )

    for _, _ = range arr {
        smallestIndex := findSmallest(arr)
        newArr = append(newArr, arr[smallestIndex])
        if smallestIndex == 0 {
            arr = append(arr[smallestIndex+1:])
        } else {
            arr = append(arr[0:smallestIndex], arr[smallestIndex+1:]...)
        }
    }

    fmt.Println("排序后的结果为: ", newArr)
}

func findSmallest(arr []int) (index int) {
    smallest := arr[0]
    smallestIndex := 0

    for i, _ := range arr {
        if arr[i] < smallest {
            smallest = arr[i]
            smallestIndex = i
        }
    }

    return smallestIndex
}
