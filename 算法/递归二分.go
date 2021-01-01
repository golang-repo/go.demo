/**
 * @Author: v_bivwei
 * @Description:
 * @Date: 2021/1/1 0:15
 */
package main

import "fmt"

const NOT_FOUND_INDEX2 = -1

func main() {
    var (
        list = []int{1, 2, 4, 6, 10, 12, 15}
    )
    index := diGuiBinarySearch(list, 0, len(list)-1, 4)
    fmt.Println("index = ", index)
}

func diGuiBinarySearch(list []int, low, high, target int) int {

    fmt.Printf("low = %d ,high = %d ,list = %+v \n", low, high, list)

    if low <= high {
        mid := (low + high) / 2 // go 会自动向下取整
        if list[mid] == target {
            return mid
        } else if list[mid] > target {
            return diGuiBinarySearch(list, 0, mid-1, target)
        } else if list[mid] < target {
            return diGuiBinarySearch(list, mid+1, high, target)
        }
    }

    return NOT_FOUND_INDEX2

}

/**

f([]int{1,2,4,6,10,12},4)





*/
