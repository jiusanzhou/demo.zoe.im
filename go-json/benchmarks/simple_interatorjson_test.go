package test

import (
	"testing"

	"github.com/json-iterator/go"
)

func Benchmark_encode_simple_interator(b *testing.B) {

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
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