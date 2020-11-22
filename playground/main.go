package main

import "fmt"

type Dp struct {
    x int
    y int
}

func main() {
    // new style for declaring
    var a = []*Dp{
        &Dp{1, 2}, 
        &Dp{3, 4}
    }
    fmt.Println(a[0].x, a[1].x)

    newStructDeclare := struct {
        x int
        y float32
    }{
        8,
        9.64,
    }
    fmt.Println(newStructDeclare);
}