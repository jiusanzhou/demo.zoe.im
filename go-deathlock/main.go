package main

import (
    "runtime"
)


func main() {
    runtime.GOMAXPROCS(1)
    
    var i int
    go func() {
        for {
            i++ 
        }
    }()

    runtime.Gosched()
    runtime.GC()
}
