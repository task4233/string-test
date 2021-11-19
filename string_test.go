package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

var args []string

const LEN = 100

// O(N^2)
func init() {
	args = make([]string, 0, LEN)
	for idx := 0; idx < LEN; idx++ {
		args = append(args, strings.Repeat("a", idx+1))
	}
}

func Benchmark(b *testing.B) {
	targets := InitForBench()

	fmt.Fprintf(os.Stderr, "len: %d\n", len(targets))
	for _, t := range targets {
		t := t

		b.Run(reflect.TypeOf(t).Elem().Name(), func(b *testing.B) {
			b.ResetTimer()
			for idx := 0; idx < b.N; idx++ {
				t.Append(args)
			}
		})
	}
}
