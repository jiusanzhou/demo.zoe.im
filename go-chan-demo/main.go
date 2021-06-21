package main

import (
    "fmt"
    "runtime"
    "time"
)

func fib(c, quit chan int) {
    x, y := 0, 1
    for {
        select {
        case c <- x:
            x, y = y, x + y
            fmt.Println("writeen")
        case <- quit:
            fmt.Println("quit")         
            return
        }
    }
}

func main() {

    runtime.GOMAXPROCS(2)

    fmt.Println("Hello, playground")
    
    c := make(chan int)
    quit := make(chan int)
    
    
    
    go func(){
        for i := 0; i < 10; i++ {
            fmt.Println("====>", <-c)
            // time.Sleep(1 * time.Millisecond)
        }
        quit <- 0
    }()

    time.Sleep(1 * time.Second)
    
    fib(c, quit)
    
}

