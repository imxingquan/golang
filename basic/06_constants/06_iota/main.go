package main

import "fmt"

const (
    _ = iota
    b = iota * 10 //1*10
    c = iota * 10 //2*10
)

func main(){
    fmt.Println(b)
    fmt.Println(c)
}