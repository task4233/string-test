package main_test

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

var errMap = map[error]int{}

const ERR_MAP_LEN = 10

func init() {
	for idx := 0; idx < ERR_MAP_LEN; idx++ {
		errMap[fmt.Errorf("error%d", idx)] = idx
	}
}

func BenchmarkMap(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		errSet := map[error]struct{}{}
		for err := range errMap {
			e := errors.Cause(err)
			if _, ok := errSet[e]; ok {
				panic("the keys of errMap do not have same root error: rootError=>" + e.Error() + ", rawError=>" + err.Error())
			}
			errSet[e] = struct{}{}
		}
	}
}

func BenchmarkMapWithCap(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		errSet := make(map[error]struct{}, len(errMap))
		for err := range errMap {
			e := errors.Cause(err)
			if _, ok := errSet[e]; ok {
				panic("the keys of errMap do not have same root error: rootError=>" + e.Error() + ", rawError=>" + err.Error())
			}
			errSet[e] = struct{}{}
		}
	}
}
