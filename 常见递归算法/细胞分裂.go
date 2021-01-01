/**
 * @Author: v_bivwei
 * @Description:
 * @Date: 2021/1/1 16:37
 */
package main

import (
    "fmt"
    "math"
)

/**
细胞分裂 有一个细胞 每一个小时分裂一次，一次分裂一个子细胞，第三个小时后会死亡。那么n个小时候有多少细胞？
*/

func main() {

    live := live(3)
    fmt.Println("live = ", live)
}

func live(hour int) int {
    if hour < 4 {
        return int(math.Pow(2, float64(hour-1)))
    } else {
        // 从第四个小时开始有死亡的细胞
        // 活着的细胞 = 前一个小时活着的细胞 - 这个小时死去的细胞
        return live(hour-1)*2 - die(hour)
    }
}

func die(hour int) int {
    if hour < 4 {
        return 0
    } else {
        /**
          因为三个小时一个周期
          也就是每三个小时,(n-3)时的细胞就会死完
          那么 这个小时(n)死去的细胞 + 上个小时(n-1)死去的细胞 + 前两个小时(n-2)死去的细胞 = 前三个小时(n-3)活着的细胞
        */
        return die(hour-3) - die(hour-2) - die(hour-1)
    }
}
