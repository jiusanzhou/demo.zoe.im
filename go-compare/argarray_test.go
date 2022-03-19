package main

import "testing"

func call_normal(x []int) {
	x[0]++
}

func call_point(x *[]int) {
	(*x)[0]++
}

func BenchmarkArgArray_Normal(b *testing.B) {
	b.ReportAllocs()

	data := make([]int, 1<<10)

	call := func(x []int) {
		x[0]++
	}

	for i := 0; i < b.N; i++ {
		call(data)
	}

}

func BenchmarkArgArray_Point(b *testing.B) {
	b.ReportAllocs()

	data := make([]int, 1<<10)

	for i := 0; i < b.N; i++ {
		call_point(&data)
	}
}
