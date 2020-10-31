Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/bionic64"
  config.vm.network "private_network", ip: "192.168.33.10"
  config.vm.synced_folder "./", "/home/vagrant/kubenet/"
  config.vm.provider "virtualbox" do |vb|
    vb.memory = "4196"
  end
  config.vm.provision "shell", inline: <<-SHELL
    add-apt-repository ppa:longsleep/golang-backports
    apt update -y
    apt install -y golang-go docker.io
  SHELL
end
