.PHONY: clean
clean:
	go clean --testcache

.PHONY: test
test:
	go test ./...

.PHONY: test-verbose
test-verbose:
	go test -v ./...

.PHONY: bench
bench:
	go test -bench ./...

.PHONY: run
run:
	go run .