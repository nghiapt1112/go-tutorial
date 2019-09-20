package main

import (
    "fmt"
    "model"
)
func main() {
    v :=model.Vertex{1,2}
    fmt.Println("original value of v: ", v)

    // to change value of v 's X property
    v.X = 123
    fmt.Println("updated value of v: ", v)

    p := v
    p.X = 1e2
    fmt.Println("pass by value", v)

    n := &v
    n.X = 1e2
    fmt.Println("pass by pointer", v)




}
