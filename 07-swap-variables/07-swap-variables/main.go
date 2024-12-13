package main

import "fmt"

func main() {
    x, y := 7, 2
    x, y = y, x
    fmt.Printf("x: %d, y: %d\n", x, y)
}