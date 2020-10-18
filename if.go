package main

import "fmt"

func main(){
    a := 10

    if a < 10 {
        fmt.Printf("a is less than 10\n")
    } else if a == 10 {
        fmt.Printf("a is %d\n", a)
    } else if a > 10 {
        fmt.Printf("a is more than 10\n", a)
    }
}
