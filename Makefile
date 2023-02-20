
OWNER := dnitsch
NAME := simplelog
GIT_TAG := "1.5.1"
VERSION := "v$(GIT_TAG)"
# VERSION := "$(shell git describe --tags --abbrev=0)"
REVISION := "aaaa11111"

.PHONY: test test_ci tidy install cross-build 

test: test_prereq
	go test `go list ./... | grep -v */generated/` -v -mod=readonly -coverprofile=.coverage/out | go-junit-report > .coverage/report-junit.xml && \
	gocov convert .coverage/out | gocov-xml > .coverage/report-cobertura.xml

test_ci:
	go test ./... -mod=readonly

test_prereq: 
	mkdir -p .coverage
	go install github.com/jstemmer/go-junit-report@v0.9.1 && \
	go install github.com/axw/gocov/gocov@v1.0.0 && \
	go install github.com/AlekSi/gocov-xml@v1.0.0

tidy: 
	go mod tidy

install: tidy
	go mod vendor

.PHONY: clean
clean:
	rm -rf bin/*
	rm -rf dist/*
	rm -rf vendor/*

tag: 
	git tag -a $(VERSION) -m "ci tag release logger" $(REVISION)
	git push origin $(VERSION)

show_coverage: test
	go tool cover -html=.coverage/out