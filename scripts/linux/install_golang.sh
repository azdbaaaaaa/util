#!/bin/bash


wget https://storage.googleapis.com/golang/go1.16.3.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.16.3.linux-amd64.tar.gz
echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/profile
source /etc/profile


# rm -rf /usr/local/go