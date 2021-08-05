package main

import "fmt"

type S struct {
	Name string
}

func (s *S) Echo() string {
	return "OK"
}

func (s *S) Validate() bool {
	fmt.Println(s.Echo())
	return s == nil
}

func run(s interface{}) {
	ss, ok := s.(*S)
	_ = ok
	fmt.Println("==>", ss.Validate())
}

func main() {
	var s *S
	fmt.Println("==>", s.Validate())
	run(nil)
}
