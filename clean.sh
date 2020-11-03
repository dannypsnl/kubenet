ip link delete kube-bridge
ip link delete nginx-1
ip link delete nginx-2
docker rm -f nginx-1 nginx-2
