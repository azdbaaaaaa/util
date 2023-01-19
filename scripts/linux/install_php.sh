#!/bin/bash

wget https://www.php.net/distributions/php-7.4.0.tar.gz
tar -zxvf php-7.4.0.tar.gz

yum install -y openssl-devel libxml2-devel bzip2-devel \
   libcurl-devel libjpeg-devel libpng-devel freetype-devel \
   libmcrypt-devel recode-devel libicu-devel libzip-devel\
   libxml2-devel sqlite-devel bzip2-devel libcurl-devel libicu-devel gcc

./configure --prefix=/usr/local/php7 --with-config-file-scan-dir=/usr/local/php7/etc/php.d --with-config-file-path=/usr/local/php7/etc --enable-mbstring --enable-zip --enable-bcmath --enable-pcntl --enable-ftp --enable-xml --enable-shmop --enable-soap --enable-intl --with-openssl --enable-exif --enable-calendar --enable-sysvmsg --enable-sysvsem --enable-sysvshm --enable-opcache --enable-fpm --enable-session --enable-sockets --enable-mbregex --enable-wddx --with-curl --with-iconv --with-gd --with-jpeg-dir=/usr --with-png-dir=/usr --with-zlib-dir=/usr --with-freetype-dir=/usr --enable-gd-jis-conv --with-openssl --with-pdo-mysql=mysqlnd --with-gettext=/usr --with-zlib=/usr --with-bz2=/usr --with-recode=/usr --with-xmlrpc --with-mysqli=mysqlnd   




sudo yum install epel-release

sudo yum install nginx


docker run --name php-fpm -p 9000:9000 -v /Users/zhoudongbin/WebstormProjects/yoshop2.0/public:/www  -d php:7.4.26-fpm
