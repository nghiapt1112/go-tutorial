package main

//func main() {
//    var a model.Abser
//    f := model.MyFloat(-math.Sqrt2)
//    v := model.Vertex{3, 4}
//
//    a = f  // a MyFloat implements Abser
//    a = &v // a *Vertex implements Abser
//
//    fmt.Println(a.Abs())
//
//    //11. Interface Values
//    fmt.Println("11. Interface Values")
//    var i model.I
//    //i.M() // nil inteface will be error cause:  there is no type inside the interface tuple to indicate which concrete method to call
//
//    // nil value
//    var t *model.T
//    i = t // nil
//    model.Describe(i)
//    i.M()  // acceptable cause: interface value that holds a nil concrete value is itself non-nil
//
//    i = &model.T{"Hello"}
//    model.Describe(i)
//    i.M()
//
//    i = model.F(math.Pi)
//    model.Describe(i)
//    i.M()
//}
