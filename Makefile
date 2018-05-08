VERSION := ${VERSION}

dep:
	go get -t ./...

build:
	go build -ldflags "-w -s -X github.com/pablo-ruth/gofat/cmd.versionString=$(VERSION)"
