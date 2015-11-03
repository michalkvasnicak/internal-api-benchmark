#!/usr/bin/env bash

apt-get install software-properties-common -y
add-apt-repository "deb http://archive.ubuntu.com/ubuntu $(lsb_release -sc) main universe"
apt-get update && apt-get install curl build-essential git pkg-config -y
curl -O https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz
curl -O -L https://github.com/nanomsg/nanomsg/releases/download/0.7-beta/nanomsg-0.7-beta.tar.gz
curl -O http://download.zeromq.org/zeromq-4.1.3.tar.gz
tar -C /usr/local -xzf go1.5.1.linux-amd64.tar.gz
tar -xzf zeromq-4.1.3.tar.gz
tar -xzf nanomsg-0.7-beta.tar.gz
cd ~/zeromq-4.1.3 && ./configure --without-libsodium && make && make install && ldconfig
cd ~/nanomsg-0.7-beta && ./configure && make && make install && ldconfig
mkdir -p ~/go; echo "export GOPATH=$HOME/go" >> ~/.bashrc
echo "export PATH=$PATH:$HOME/go/bin:/usr/local/go/bin" >> ~/.bashrc
source ~/.bashrc
go get github.com/michalkvasnicak/internal-api-benchmark