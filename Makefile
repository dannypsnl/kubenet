.PHONY: run clean
run:
	@go build
	@./kubenet
clean:
	@ip link delete kube-bridge
	@ip link delete kni0
