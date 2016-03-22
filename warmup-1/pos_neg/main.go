package main

import "fmt"

func pos_neg(a int, b int, negative bool) bool {
    if negative {
        return (a < 0 && b < 0)
    } else {
        return ((a < 0 && b > 0) || (a > 0 && b < 0))
    }
}

func main(){
    fmt.Println(pos_neg(1,-1,false))
    fmt.Println(pos_neg(-1,1,false))
    fmt.Println(pos_neg(-4,-5,true))
    fmt.Println(pos_neg(-4,-5,false))
    fmt.Println(pos_neg(-4,5,false))
    fmt.Println(pos_neg(-4,5,true))
    fmt.Println(pos_neg(1,1,false))
    fmt.Println(pos_neg(-1,-1,false))
    fmt.Println(pos_neg(1,-1,true))
    fmt.Println(pos_neg(-1,1,true))
    fmt.Println(pos_neg(1,1,true))
    fmt.Println(pos_neg(-1,-1,true))
    fmt.Println(pos_neg(5,-5,false))
    fmt.Println(pos_neg(-6,6,false))
    fmt.Println(pos_neg(-5,-6,false))
    fmt.Println(pos_neg(-2,-1,false))
    fmt.Println(pos_neg(1,2,false))
    fmt.Println(pos_neg(-5,6,true))
    fmt.Println(pos_neg(-5,-5,true))
}
