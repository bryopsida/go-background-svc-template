clean:
	rm -rf bin/*
build:
	go build -o bin/service main.go

image:
	docker build -t ghcr.io/bryopsida/go-background-svc-template:local .
