package main

import (
    "fmt"
)

type Vertex struct {
    Lat, Long float64
}

func myMap() {
    var m map[string]Vertex
    m = make(map[string]Vertex)

    m["key1"] = Vertex{1, -1}
    m["key2"] = Vertex{12, -12}
    fmt.Println(m["key1"])
    fmt.Println(m["key2"])

}

/**
Map literals are like struct literals, but the keys are required.
*/
func mapLiterals() {
    var m = map[string]Vertex{
        "Bell Labs": {1, -2,},
        "Google":    {3, -3,},
    }
    fmt.Println(m)
}

func mutatingMap() {
    myMap := make(map[string]int)

    myMap["key1"] = 11
    myMap["key2"] = 12

    fmt.Println("gia tri cua map la: ", myMap)
    fmt.Println("gia tri cua key1 la: ", myMap["key1"])

    fmt.Println("gia tri cua key2 la: ", myMap["key2"])
    fmt.Println("===== deleting key1 ")
    delete(myMap, "key1")
    fmt.Println("gia tri cua key1 la: ", myMap["key1"])
    fmt.Println("===== check key1 is existed?")

    value, existed := myMap["key1"]
    fmt.Println("The value:", value, "Present?", existed)

}

func WordCount(s string) {
    fmt.Println(map[string]int{"x": 1})
}

func main() {
    //myMap()
    //mapLiterals()
    mutatingMap()


}
