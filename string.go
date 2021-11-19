package main

import (
	"strings"
)

type T interface {
	Append([]string) string
}

func InitForBench() []T {
	return []T{
		&WithOperator{},
		&WithJoin{},
		&WithBytes{},
		&WithBytesSetCap{},
	}
}

type WithOperator struct{}

func (w *WithOperator) Append(args []string) (res string) {
	for idx := range args {
		res += args[idx]
	}
	return
}

type WithJoin struct{}

func (w *WithJoin) Append(args []string) (res string) {
	res = strings.Join(args, "")
	return
}

type WithBytes struct{}

func (w *WithBytes) Append(args []string) (res string) {
	b := []byte{}
	for idx := range args {
		b = append(b, args[idx]...)
	}
	res = string(b)
	return
}

type WithBytesSetCap struct{}

func (w *WithBytesSetCap) Append(args []string) (res string) {
	// 今回は、argsに含まれる文字列長を1, 2, ..., nという問題設定とするので、(1+n)*n/2あれば良さそう
	b := make([]byte, 0, (1+len(args))*len(args)/2)
	for idx := range args {
		b = append(b, args[idx]...)
	}
	res = string(b)
	return
}
