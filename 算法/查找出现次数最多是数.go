/**
 * @Author: v_bivwei
 * @Description:
 * @Date: 2020/12/28 22:23
 */
package main

import "fmt"

/**
  [1,2,3,4,5,5,5,6] = 5 出现了 3 次,输出 5
*/
func main() {
    var (
        s       = []int{1, 2, 3, 4, 5, 5, 5, 6}
        sMap    = make(map[int]int, 0)
        valMax  = -1
        valTime = 0
    )

    for _, v := range s {
        if _, ok := sMap[v]; ok {
            sMap[v] ++
        } else {
            sMap[v] = 1
        }
    }

    for sVal, time := range sMap {
        if time > valTime {
            valTime = time
            valMax = sVal
        }
    }

    fmt.Println(sMap, valMax, valTime)

}
