runbench:
	go test -bench=. -v -benchmem  > bench.out
protogo:
	protoc *.proto -I. -I=$(GOPATH)/src --go_out=plugins=grpc:$(GOPATH)/src