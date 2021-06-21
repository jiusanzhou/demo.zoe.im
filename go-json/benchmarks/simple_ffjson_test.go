package test

import (
	"testing"

	"github.com/pquerna/ffjson/ffjson"
)

func Benchmark_encode_simple_ffjson(b *testing.B) {
	var l int64
	for i := 0; i < b.N; i++ {
		data, err := ffjson.Marshal(&zoe)
		if err != nil {
			b.Error(err)
		}
		l = int64(len(data))
	}
	b.SetBytes(l)
}