build:
	go build -o ricer ./cmd/ricer
run: build
	./ricer
