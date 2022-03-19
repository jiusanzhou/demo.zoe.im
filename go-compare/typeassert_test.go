package main

import (
	"testing"
)

func BenchmarkTypeAssert_Noop(b *testing.B) {
	b.ReportAllocs()

	call := func(p Person) {
		pp := p
		pp.Name += ""
	}

	me := Person{Name: "Zoe"}
	for i := 0; i < b.N; i++ {
		call(me)
	}
}

func BenchmarkTypeAssert_Specific(b *testing.B) {
	b.ReportAllocs()

	call := func(p interface{}) {
		pp, _ := p.(Person)
		pp.Name += ""
	}

	me := Person{Name: "Zoe"}
	for i := 0; i < b.N; i++ {
		call(me)
	}
}
