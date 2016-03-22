package main

import "fmt"

func monkey_trouble(a_smile bool, b_smile bool) bool {
    if (a_smile && b_smile) || (!a_smile && !b_smile) {
        return true
    } else {
        return false
    }
}

func main(){
    fmt.Println(monkey_trouble(true,true))
    fmt.Println(monkey_trouble(false,false))
    fmt.Println(monkey_trouble(true,false))
    fmt.Println(monkey_trouble(false,true))


}
