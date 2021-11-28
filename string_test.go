package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

var args []string

// O(N^2)
func init() {
	LEN, err := strconv.Atoi(os.Getenv("ARG_LEN"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "ARG_LEN is not set")
		os.Exit(1)
	}

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
