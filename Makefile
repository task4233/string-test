.PHONY: bench
bench:
	go test -bench . -benchmem -count=3
