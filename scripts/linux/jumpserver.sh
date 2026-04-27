#! /usr/bin/expect

set passwd jzOrz8_Vb52SHLdR

spawn ssh zhoudongbin@172.24.0.16 -p60022 -i /Users/zhoudongbin/.ssh/zhoudongbin.pem
expect "*Enter passphrase*"
send "${passwd}\n"
interact
