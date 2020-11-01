.PHONY: run clean
run:
	@go build ./cmd/kubenet/
	@./kubenet
clean:
	@ip link delete kube-bridge
	@ip link delete kni0
