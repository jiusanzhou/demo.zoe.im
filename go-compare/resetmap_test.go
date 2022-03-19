package main

import "testing"

func BenchmarkResetMap_Delete(b *testing.B) {
	b.ReportAllocs()
	size := 1 << 10
	data := make(map[int]int, size)

	for i := 0; i < b.N; i++ {
		// inset data
		for i := 0; i < size; i++ {
			data[i] = i
		}

		// delete
		for k := range data {
			delete(data, k)
		}
	}
}

func BenchmarkResetMap_Remake(b *testing.B) {
	b.ReportAllocs()
	size := 1 << 10
	data := make(map[int]int, size)

	for i := 0; i < b.N; i++ {

		// inset data
		for i := 0; i < size; i++ {
			data[i] = i
		}

		data = make(map[int]int, size)
	}
}
