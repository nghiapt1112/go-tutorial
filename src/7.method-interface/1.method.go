package main

import (
    "math"
    "model"
)

func Abs(v model.Vertex) float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func ScaleFunc(v *model.Vertex, f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

//func main() {
//    v := model.Vertex{3, 4}
//    fmt.Println("Value before scale: ")
//    fmt.Println(v.Abs())
//    v.Scale(10)
//    fmt.Println("Value after scale: ")
//    fmt.Println(v.Abs())
//    fmt.Println(Abs(v))
//
//    fmt.Println("ScaleFunc by function in other package: ")
//    ScaleFunc(&v, 10)
//
//    fmt.Println(v.Abs())
//    fmt.Println(Abs(v))
//
//    fmt.Println("Print an object and print a pointer: ")
//    fmt.Println(v, &v)
//}
