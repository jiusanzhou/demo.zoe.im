package main

import (
	"fmt"
	"testing"
)

func BenchmarkStringConcat_Plus(b *testing.B) {
	b.ReportAllocs()
	me := Person{Name: "Zoe", Age: 42}
	for i := 0; i < b.N; i++ {
		_ = me.Name + me.String() + "abcdefg"
	}
}

func BenchmarkStringConcat_Sprinf(b *testing.B) {
	b.ReportAllocs()
	me := Person{Name: "Zoe", Age: 42}
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s%s", me.Name, me.String(), "abcdefg")
	}
}
