package main

import "fmt"

func level3(x int){
    x = x+1
    fmt.Printf("level3: %d\n", x)
}

func level2(x int){
    x = x+1
    fmt.Printf("level2: %d\n", x)
    level3(x)
}

func level1(x int){
    x = x+1
    fmt.Printf("level1: %d\n", x)
    level2(x)
}

func main(){
    a := 1
    fmt.Printf("start: %d\n", a)
    level1(a)
}
