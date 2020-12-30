/**
 * @Author: v_bivwei
 * @Description:
 * @Date: 2020/12/30 21:09
 */
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

func main() {
    log.Println("121111")
    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        log.Println("3333")
        s := fmt.Sprintf("世界,你好.time %s ", time.Now().String())
        _, _ = fmt.Fprintf(writer, "%s \n", s)
        log.Printf("start %s \n", s)
    })

    err := http.ListenAndServe(":8082", nil)
    if err != nil {
        log.Fatalf("listen error: %s \n", err.Error())
    }

}
