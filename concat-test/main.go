package main

func concatWithConstString(n int) string {
	var s string
	const constHoge = "hoge"
	for idx := 0; idx < n; idx++ {
		s += constHoge
	}
	return s
}

func concatWithVarString(n int) string {
	var s string
	var varHoge = "hoge"
	for idx := 0; idx < n; idx++ {
		s += varHoge
	}
	return s
}

func concatWithRawString(n int) string {
	var s string
	for idx := 0; idx < n; idx++ {
		s += "hoge"
	}
	return s
}

func main() {
	concatWithConstString(100)
	concatWithVarString(100)
	concatWithRawString(100)
}
