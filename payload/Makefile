runbench:
	go test -bench=. -v -benchmem #-memprofile=memAll.out
pprof:
	go tool pprof memAll.out
protogo:
	protoc *.proto -I. -I=$(GOPATH)/src --go_out=plugins=grpc:$(GOPATH)/src
