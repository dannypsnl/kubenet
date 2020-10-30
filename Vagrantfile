Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/bionic64"
  config.vm.network "private_network", ip: "192.168.33.10"
  config.vm.synced_folder "./", "/home/vagrant/kube/"
  config.vm.provider "virtualbox" do |vb|
      vb.memory = "4196"
  end
end
