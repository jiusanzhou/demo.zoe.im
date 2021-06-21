package main

import (
    "fmt"
    "io/fs"

    _ "embed"
)

///go:embed test.txt
var s string

func main() {
    fmt.Println(fs.ValidPath("test.txt"))
    fmt.Println("==>", s)
}
