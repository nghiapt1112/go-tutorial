package utils

import (
    "fmt"
    "model"
)
func describe(i model.I) {
    fmt.Printf("(%v, %T)\n", i, i)
}