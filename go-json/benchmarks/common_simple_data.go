package test

//go:generate ffjson common_simple_data.go

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Male bool   `json:"male"`
}

var zoe = Person{ "Zoe", 26, true }