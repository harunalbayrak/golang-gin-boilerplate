.PHONY: all
all:
	go build -o ./build/golang-gin-boilerplate ./cmd/main.go

.PHONY: docker
docker:
	docker build -t harunalbayrak/golang-gin-boilerplate:latest .

.PHONY: run
run:
	go run ./cmd