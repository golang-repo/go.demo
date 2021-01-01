/**
 * @Author: v_bivwei
 * @Description:
 * @Date: 2021/1/1 0:55
 */
package main

import "fmt"

func main() {
    greet("吴彦祖")
}

func greet2(name string) {
    fmt.Printf("how are you %s \n", name)

}

func bye() {
    fmt.Println("ok ,bye")
}

func greet(name string) {
    fmt.Println("hello ", name, " !")
    greet2(name)
    bye()
}
