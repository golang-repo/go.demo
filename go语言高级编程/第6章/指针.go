/**
 * @Author: v_bivwei
 * @Description:
 * @Date: 2021/1/2 17:54
 */

package main

import "fmt"

func double(x *int) {
    *x += *x
}
func main() {
    var a = 3
    double(&a)
    fmt.Println(a) // 6
    p := &a
    double(p)
    fmt.Println(a, p == nil) // 12 false
}
