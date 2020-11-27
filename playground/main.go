package main

import (
    "fmt"
    "time"
)

// Dp struct
type Dp struct {
    x int
    y int
}

func main() {
    // new style for declaring
    var a = []*Dp{
        &Dp{1, 2}, 
        &Dp{3, 4},
    }
    fmt.Println(a[0].x, a[1].x)

    newStructDeclare := struct {
        x int
        y float32
    }{
        8,
        9.64,
    }
    fmt.Println(newStructDeclare)

    // func monthDayYear(t time.Time) string {
    //     return t.Format("00-00-0000")
    // }
    fmt.Println(time.Now())
    fmt.Println(time.Now().Format(time.RFC822))
    fmt.Println(time.Now().Format("02-01-2006"))

    fmt.Println("this is from remote docker container")
}