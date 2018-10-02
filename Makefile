test:
	go test -v $(go list ./... | grep -v /vendor/)

race:
	go test -race -v $(go list ./... | grep -v /vendor/)	