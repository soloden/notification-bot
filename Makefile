.PHONY:

build:
	go build -o ./bin/notificator ./cmd/notificator/main.go

run: build
	./bin/notificator