package test

import (
	"encoding/json"
	"testing"
)

func Benchmark_encode_simple_std(b *testing.B) {
	var l int64
	for i := 0; i < b.N; i++ {
		data, err := json.Marshal(&zoe)
		if err != nil {
			b.Error(err)
		}
		l = int64(len(data))
	}
	b.SetBytes(l)
}