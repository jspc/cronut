default: build

build:
	go build

linux:
	GOOS=linux go build -o cronut-linux

 dist: linux
	docker build -t jspc/cronut .

push: dist
	docker push jspc/cronut
