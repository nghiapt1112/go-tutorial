package main

import "fmt"

/**
A var declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
*/
var a, b, c = 1, 2, 3
var i, j int = 1, 2

/**
:= nếu đc khai báo trong func thì nó sẽ tương đường vs truyền giá trị và gán kiểu, nhưng không dùng đc ngoài func
*/
func main() {
    var d, e, f = true, 2, " string day"
    fmt.Println(a, b, c, d, e, f)

    k := 3
    c, python, java := true, false, "no!"
    fmt.Println(i, j, k, c, python, java)

}
