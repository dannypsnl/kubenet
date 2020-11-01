SRC := $(shell ls *.go)

.PHONY: build run clean
build: $(SRC)
	@go build ./cmd/kubenet/
run: build
	@./kubenet
clean:
	@ip link delete kube-bridge
	@ip link delete nginx-1
	@ip link delete nginx-2
