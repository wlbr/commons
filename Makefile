LINKERFLAGS = -X main.Version=`git describe --tags --always --long --dirty` -X main.BuildTimestamp=`date -u '+%Y-%m-%d_%I:%M:%S_UTC'`
PROJECTROOT = $(dir $(abspath $(lastword $(MAKEFILE_LIST))))


all: clean build

.PHONY: clean
clean:
	@echo Running clean job...
	rm -rf bin/
	rm -f main
	@##if not checking in the generated source, then it should be deleted during the clean task
	@#rm -f logger/loglevel_string.go
	rm -f coverage.txt


logger/loglevel_enumer.go: logger/loglevel.go
	go generate logger/loglevel.go

generate: dep logger/loglevel_enumer.go
	@echo Running generate job...


build: generate
	@echo Running build job...
	@#mkdir -p bin/linux bin/windows bin/mac
	GOOS=linux go build  -ldflags "$(LINKERFLAGS)"  ./...
	GOOS=windows go build  -ldflags "$(LINKERFLAGS)" ./...
	GOOS=darwin go build  -ldflags "$(LINKERFLAGS)"  ./...


test: generate
	@echo Running test job...
	go test ./... -cover -coverprofile=coverage.txt

coverage: test
	@echo Running coverage job...
	go tool cover -html=coverage.txt



$(GOPATH)/src/github.com/alvaroloes/enumer:
	go get -u github.com/alvaroloes/enumer


dep: $(GOPATH)/src/github.com/alvaroloes/enumer
	@echo Running dep job...

