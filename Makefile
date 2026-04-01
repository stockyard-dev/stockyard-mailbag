build:
	CGO_ENABLED=0 go build -o mailbag ./cmd/mailbag/

run: build
	./mailbag

test:
	go test ./...

clean:
	rm -f mailbag

.PHONY: build run test clean
