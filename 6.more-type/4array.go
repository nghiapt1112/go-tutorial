package main

import "fmt"

func myArray() {
    var a [2] string
    a[0] = "Nghia"
    a[1] = "Pham"

    fmt.Println(a[0], a[1])
    fmt.Println(a)

    b := [6] int{1, 2, 3, 4, 5, 5}
    fmt.Println(b)
}

/**
A slice does not store any data, it just describes a section of an underlying array.
Changing the elements of a slice modifies the corresponding elements of its underlying array.
Other slices that share the same underlying array will see those changes.


khi slice se khong tao ra mang moi, no chi cat cai mang cu ra va van giu reference den mang cu do => thay doi gia tri cua mang slice thi se thay doi gia tri cua mang cu

slice cung tao ra 1 mang => co day du tinh nang cua 1 mang.
*/
func slices(start, end int) {
    fmt.Println("\n slice:")
    a := [8] string{"value1", "value2", "value3", "value4", "value5", "value6", "value7", "value8"}
    if start > end {
        fmt.Println("Gia tri khong phu hop, start must be less than end")
    } else {
        var s = a[start:end]
        fmt.Printf("mang sau khi slice %v ->  %v: %v\n", start, end, s)
        s[2] = "HEHEHEHE"
        fmt.Printf("mang sau slice khi update value1 %v\n", s)
        fmt.Printf("mang origin khi update value1 %v\n", a)
    }

}

func nilCheck() {
    var s []int
    fmt.Println(s, len(s), cap(s))
    if s == nil {
        fmt.Println("nil!")
    }
}

func shortSlice() {
    fmt.Println("\nshort slice:")
    s := []int{2, 3, 5, 7, 11, 13}

    a := s[1:4]
    fmt.Println(a)

    b := s[:2] // equivalent: s[0:2]
    fmt.Println(b)

    c := s[2:] // equivalent: s[2:len(s)]
    fmt.Println(c)
}

func main() {
    myArray()
    slices(1, 4)
    shortSlice()
    nilCheck()
}
