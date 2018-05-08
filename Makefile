VERSION := ${VERSION}

build:
	go build -ldflags "-w -s -X github.com/pablo-ruth/gofat/cmd.versionString=$(VERSION)"
