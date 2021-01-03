/**
 * @Author: v_bivwei
 * @Description:
 * @Date: 2021/1/3 21:19
 */
package main

import "fmt"

/**
一个数列：1,1,2,3,5,8... 指从第三个值开始，每个值都是前两个值的和
*/
func main() {
    ret := fbnq(6)
    fmt.Printf("结果为: %d \n", ret)
}

func fbnq(num int) int {
    if num < 3 {
        return 1
    }

    return fbnq(num-1) + fbnq(num-2)

}
