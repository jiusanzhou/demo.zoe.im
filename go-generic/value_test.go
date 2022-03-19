package main

import (
	"testing"

	"go.zoe.im/x"
)

func BenchmarkXValue_Normal(b *testing.B) {
	b.ReportAllocs()

	call := func(x int) int {
		if x == 0 {
			return -1
		} else {
			return x
		}
	}

	for i := 0; i < b.N; i++ {
		call(0)
		call(1)
		call(0)
		call(1)
	}
}

func BenchmarkXValue_InterfaceCall(b *testing.B) {
	b.ReportAllocs()
	call := func(v int) (int, bool) {
		return x.NewValue(v).Or(-1).Int()
	}

	for i := 0; i < b.N; i++ {
		call(0)
		call(1)
		call(0)
		call(1)
	}
}

func BenchmarkXValue_GenericCall(b *testing.B) {
	b.ReportAllocs()
	call := func(v int) int {
		return NewValue(v).Or(-1).Value()
	}

	for i := 0; i < b.N; i++ {
		call(0)
		call(1)
		call(0)
		call(1)
	}
}
