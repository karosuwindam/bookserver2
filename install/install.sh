#!/usr/bin/bash

go_version=1.18

curl -OL https://go.dev/dl/go$go_version.linux-arm64.tar.gz &&\
sudo tar -C /usr/local -xzf go$go_version.linux-arm64.tar.gz &&\
rm -rf go$go_version.linux-arm64.tar.gz
echo "export PATH=$PATH:/usr/local/go/bin" >> $HOME/.bashrc &&\
source $HOME/.bashrc
