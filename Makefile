include variables.mk

run:
	$(GORUN) main.go -filename stage/sample-small.csv -appid 34

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(build) .

test:
	$(GOTEST) ./...

.PHONY: run