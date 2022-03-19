package main

import "testing"

func BenchmarkResetSlice_Slice(b *testing.B) {
	b.ReportAllocs()
	size := 1 << 10
	data := make([]int, size)
	temp := make([]int, size)
	// inset data
	for i := 0; i < size; i++ {
		temp[i] = i
	}

	for i := 0; i < b.N; i++ {
		copy(data, temp)
		data = data[:]
	}
}

func BenchmarkResetSlice_Setnil(b *testing.B) {
	b.ReportAllocs()
	size := 1 << 10
	data := make([]int, size)
	temp := make([]int, size)
	// inset data
	for i := 0; i < size; i++ {
		temp[i] = i
	}

	for i := 0; i < b.N; i++ {
		copy(data, temp)
		data = nil
	}
}

func BenchmarkResetSliceAppend_Slice(b *testing.B) {
	b.ReportAllocs()
	size := 1 << 10
	data := make([]int, size)

	for i := 0; i < b.N; i++ {
		// inset data
		for i := 0; i < size; i++ {
			data = append(data, i)
		}
		data = data[:]
	}
}

func BenchmarkResetSliceAppend_Setnil(b *testing.B) {
	b.ReportAllocs()
	size := 1 << 10
	data := make([]int, size)

	for i := 0; i < b.N; i++ {
		// inset data
		for i := 0; i < size; i++ {
			data = append(data, i)
		}
		data = nil
	}
}
