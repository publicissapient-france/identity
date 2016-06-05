all: compile build-image

compile:
	docker run --rm -v $(PWD):/usr/src/identity -w /usr/src/identity \
     			 	-e CGO_ENABLED=0 -e GOOS=linux -e GOARCH=amd64 \
     			 	golang:1.6 go get -v -d && go build -v

build-image:
	docker build -t identity .