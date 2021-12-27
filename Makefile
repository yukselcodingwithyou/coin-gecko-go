build:
	go build -o main

run:
	./main $(option)

ping:
	curl http://localhost:1234/ping