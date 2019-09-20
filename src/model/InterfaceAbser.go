package model

import "fmt"

//
//Interface
//
type Abser interface {
    Abs() float64
}

type I interface {
    M()
}

func Describe(i I) {
    fmt.Printf("(%v, %T)\n", i, i)
}