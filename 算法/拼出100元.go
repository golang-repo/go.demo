/**
 * @Author: v_bivwei
 * @Description:
 * @Date: 2020/12/28 22:14
 */
package main

import "fmt"

/**
用 2 , 3 ,7 拼出 100 元,有多少种可能性
*/

func main() {
    var count int = 1
    fmt.Println(count, 100/7)

    for i := 0; i < 100/7; i++ {
        for j := 0; j < 100/3; j++ {
            if (100-7*i-3*j)%2 == 0 {
                count++
            }
        }
    }

    fmt.Println("count= ", count)

}
