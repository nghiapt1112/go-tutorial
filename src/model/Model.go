package model

import (
    "fmt"
    "math"
)

//
//Object1 implement Abser by auto implement method Abs()
//
type Vertex struct {
    X, Y float64
}

// this method belong to Vertex type.
func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// pointer = reference on JAVA
func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

//
type T struct {
    S string
}
func (t *T) M() {
    if t == nil {
        fmt.Println("<nil>")
        return
    }
    fmt.Println(t.S)
}

type F float64

func (f F) M() {
    fmt.Println(f)
}