roller: clean
	go build server.go roller.go handlers.go

test:
	go test server.go roller.go handlers.go
	
clean:
	rm -f server