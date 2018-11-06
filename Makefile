NAME		:= birr
VERSION	:= v0.0.1

.PHONY: all
all: install build

.PHONY: install
install:
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get google.golang.org/grpc

.PHONY: build
build: birrd birr

.PHONY: birrd
birrd: protoc
	go build -o bin/birrd birrd/main.go

.PHONY: birr
birr: protoc
	go build -o bin/birr birr/main.go

.PHONY: protoc
protoc:
	protoc -I api/ --go_out=plugins=grpc:api api/api.proto

.PHONY: clean
clean:
	rm bin/birrd
	rm bin/birr
#	rm api/api.pb.go
