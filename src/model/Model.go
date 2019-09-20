package model

import "math"

type Vertex struct {
    X, Y float64
}
// this method belong to Vertex type.
func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// pointer = reference on JAVA
func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}