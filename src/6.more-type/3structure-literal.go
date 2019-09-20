package main

import (
    "fmt"
    "model"
)

var (
    v1 = model.Vertex{1, 2}  // has type Vertex
    v2 = model.Vertex{X: 1}  // Y:0 is implicit
    v3 = model.Vertex{}      // X:0 and Y:0
    p  = &model.Vertex{1, 2} // has type *Vertex
)

func main() {
    fmt.Println(v1, p, v2, v3)
}
