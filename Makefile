.PHONY: test
test:
	go test -v $(go list ./... | grep -v /vendor/)

race:
	go test -race -v $(go list ./... | grep -v /vendor/)

dep:
	glide install

fmt:
	gofmt -w .

lint:
	./scripts/pre-commit.sh

ci: dep fmt lint race
