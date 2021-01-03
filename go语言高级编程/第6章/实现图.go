/**
 * @Author: v_bivwei
 * @Description:
 * @Date: 2021/1/2 10:29
 */
package main

import "fmt"

func main() {

    var ()

    graph := getGraph()
    ret := searchV1(graph, "you")
    fmt.Printf("结果为 %t ", ret)
}

// 以 m 结尾就是供应商
func isSeller(person string) bool {
    return person[len(person)-1:len(person)] == "m"
}

func searchV1(graph map[string][]string, name string) bool {
    var (
        searchQueue = make([]string, 0)
        searched    = make([]string, 0)
    )

    searchQueue = append(searchQueue, graph[name]...)

    for len(searchQueue) > 0 {
        person := searchQueue[0] // 要查找的人
        searchQueue = searchQueue[1:]
        personAlreadySearched := false

        // 看是否已经查过了
        for _, val := range searched {
            if val == person {
                personAlreadySearched = true
            }
        }

        if personAlreadySearched == false {
            if isSeller(person) {
                fmt.Printf("找个了供应商： %s \n", person)
                fmt.Println("search === ", searched)
                return true
            }

            searched = append(searched, person) // 把这个人放到已搜索中
            searchQueue = append(searchQueue, graph[person]...)
        }

    }

    return false

}

func getGraph() (graph map[string][]string) {
    graph = make(map[string][]string)

    graph["you"] = []string{"alice", "bob", "claire"}
    graph["bob"] = []string{"anuj", "peggy"}
    graph["alice"] = []string{"peggy"}
    graph["claire"] = []string{"thom", "jonny"}
    graph["anuj"] = []string{}
    graph["peggy"] = []string{}
    graph["thom"] = []string{}
    graph["jonny"] = []string{}

    return graph

}
