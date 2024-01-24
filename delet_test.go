package main

import "testing"

func BenchmarkCoreV1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := make([]int, 128)
		for i := 0; i < 20; i++ {
			CoreV1(s, i)
		}
	}
}

func BenchmarkCoreV2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := make([]int, 128)
		for i := 0; i < 20; i++ {
			CoreV2(s, i)
		}
	}
}
