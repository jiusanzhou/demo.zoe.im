package test

// import (
// 	"testing"

// 	"github.com/mailru/easyjson"
// )

// func Benchmark_encode_simple_easyjson(b *testing.B) {
// 	var l int64
// 	for i := 0; i < b.N; i++ {
// 		data, err := easyjson.Marshal(&zoe)
// 		if err != nil {
// 			b.Error(err)
// 		}
// 		l = int64(len(data))
// 	}
// 	b.SetBytes(l)
// }