package main

import "fmt"
import "time"
import "math/rand"

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, " : ", i)
    }
}

func main() {
    f("direct")

    go f("goroutine")

    go func(msg string) {
        fmt.Println(msg)
    }("going")
    fmt.Scanln()
    fmt.Println("Done")

    msgs := make(chan int)
    for i := 0; i < 4; i++{
        go func() {
            msgs <- rand.Int()
            time.Sleep(2 * time.Second)
        }()
    }

    for res := range msgs {
        fmt.Println(res)
    }
}
