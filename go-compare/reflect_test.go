package main

import (
	"reflect"
	"testing"
)

func BenchmarkReflect_Noop(b *testing.B) {
	b.ReportAllocs()

	me := Person{Name: "Zoe", Age: 30}
	var r string
	for i := 0; i < b.N; i++ {
		r = me.String()
	}
	_ = r
}

func BenchmarkReflect_Call(b *testing.B) {
	b.ReportAllocs()

	me := Person{Name: "Zoe", Age: 30}
	var r string
	for i := 0; i < b.N; i++ {
		r = reflect.ValueOf(me).MethodByName("String").
			Call(nil)[0].
			Interface().(string)
	}
	_ = r
}
