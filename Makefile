SRC := $(shell ls *.go)

.PHONY: build run clean
build: $(SRC)
	@go build ./cmd/kubenet/
run: build
	@./kubenet
clean:
	@./clean.sh
