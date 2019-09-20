package main

import (
    "fmt"
    "model"
)

func main() {
    a := model.Person{"Arthur Dent", 42}
    z := model.Person{"Zaphod Beeblebrox", 9001}
    fmt.Println(a, z)
}
