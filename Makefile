default: build

build:
	go build

linux:
	docker run -v $(PWD):/root/golang/src/github.com/jspc/cronut -v $(GOPATH)/src/github.com/zeebox:/root/golang/src/github.com/zeebox --workdir /root/golang/src/github.com/jspc/cronut jspc/alpine-build

dist: linux
	docker build -t jspc/cronut .

push: dist
	docker push jspc/cronut
