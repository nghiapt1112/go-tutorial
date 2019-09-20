package main

import "fmt"

/**
The make function allocates a zeroed array and returns a slice that refers to that array:
*/

func main() {
    a := make([]int, 5)
    printSlice("a", a)

    b := make([]int, 0, 5)
    printSlice("b", b)


    c := b[:2]
    printSlice("c", c)

    d := c[2:5]
    printSlice("d", d)


    fmt.Println("-------------append slice------------------")
    appendSlice()
}

func appendSlice() {
    var s []int
    printSlice("",s)

    // append works on nil slices.
    s = append(s, 0)
    printSlice("",s)

    // The slice grows as needed.
    s = append(s, 1)
    printSlice("",s)

    // We can add more than one element at a time.
    s = append(s, 2, 3, 4)
    printSlice("",s)
}


func printSlice(s string, x []int) {
    fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
