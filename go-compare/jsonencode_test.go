package main

import (
	"bytes"
	"encoding/json"
	"testing"
)

func BenchmarkEncode_Normal(b *testing.B) {
	b.ReportAllocs()
	var buf []byte
	me := Person{
		Name: "John", Age: 42, Email: "zoe@example.com",
	}
	for i := 0; i < b.N; i++ {
		buf, _ = json.Marshal(me)
		buf = nil
	}
	_ = buf
}

func BenchmarkEncode_Encoder(b *testing.B) {
	b.ReportAllocs()
	var buf []byte
	me := Person{
		Name: "John", Age: 42, Email: "zoe@example.com",
	}
	x := json.NewEncoder(bytes.NewBuffer(buf))
	for i := 0; i < b.N; i++ {
		x.Encode(me)
		buf = nil
	}
}
