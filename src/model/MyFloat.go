package model

//
//MyFloat implement Abser by auto implement method Abs()
//
//
type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}
