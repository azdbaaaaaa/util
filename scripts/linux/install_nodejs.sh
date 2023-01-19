#!/bin/bash


curl --silent --location https://rpm.nodesource.com/setup_8.x | sudo bash -
sudo yum -y install nodejs

npm i n -g
n v12.16.3
