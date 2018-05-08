VERSION := ${VERSION}

build:
	GOOS=linux GOARCH=amd64 go build -o gofat-linux-amd64 -ldflags "-w -s -X github.com/pablo-ruth/gofat/cmd.versionString=$(VERSION)"
	GOOS=linux GOARCH=386 go build -o gofat-linux-386 -ldflags "-w -s -X github.com/pablo-ruth/gofat/cmd.versionString=$(VERSION)"

clean:
	rm gofat-linux-amd64 gofat-linux-386

dep:
	go get -t ./...
