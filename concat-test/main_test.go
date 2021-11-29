package main

import (
	"os"
	"strconv"
	"testing"
)

var N int

func init() {
	if os.Getenv("ARG_N") == "" {
		panic("ARG_N must not be empty")
	}

	var err error
	N, err = strconv.Atoi(os.Getenv("ARG_N"))
	if err != nil {
		panic("failed to convert")
	}
}

func BenchmarkConcatWithConstString(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		concatWithConstString(N)
	}
}

func BenchmarkConcatWithVarString(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		concatWithVarString(N)
	}
}

func BenchmarkConcatWithRawString(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		concatWithRawString(N)
	}
}
