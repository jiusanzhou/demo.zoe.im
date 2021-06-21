package main

import (
	"fmt"
	"os/exec"

	"go.zoe.im/x"
)

func main() {

	fmt.Println("Please press Ctrl+C ...")

	x.GraceRun(func() error {

		return nil
	})



	// Command * is not work should be `bash -c ""`
	exec.Command("sleep", "5").Run()
	exec.Command("mkdir", "demo1").Run()

	fmt.Println("Exit!")
}