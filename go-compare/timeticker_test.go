package main

import (
	"testing"
	"time"
)

func BenchmarkTimeTicker_Normal(b *testing.B) {
	b.ReportAllocs()
	i := 0
	t := time.NewTicker(time.Millisecond)
	for {
		<-t.C
		i++
		if i > b.N {
			break
		}
	}
}

func BenchmarkTimeTicker_Range(b *testing.B) {
	b.ReportAllocs()
	i := 0
	for range time.Tick(time.Millisecond) {
		i++
		if i > b.N {
			break
		}
	}
}
