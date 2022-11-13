roller: clean
	go build server.go roller.go handlers.go

test:
	go test
	
clean:
	rm server