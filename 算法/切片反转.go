/**
 * @Author: v_bivwei
 * @Description:
 * @Date: 2020/12/28 21:51
 */

/**
  输入一个数组 [1,2,3,4] 输出 [4,3,2,1]
*/
package main

import "fmt"

func main() {
    var (
        dst = []int{1, 2, 3, 4, 5}
    )
    src := reverse01(dst)
    fmt.Println(src)
    fmt.Println(reverse02(dst))
}

func reverse01(dst []int) (src []int) {
    for i := len(dst) - 1; i >= 0; i-- {
        src = append(src, dst[i])
    }

    return src
}

func reverse02(dst []int) (src []int) {
    tmp := 0
    for i := 0; i < len(dst)/2; i++ {
        tmp = dst[len(dst)-i-1]
        dst[len(dst)-i-1] = dst[i]
        dst[i] = tmp
    }

    return dst
}
