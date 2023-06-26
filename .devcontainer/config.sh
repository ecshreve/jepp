#!/bin/sh

set -e

export DEBIAN_FRONTEND=noninteractive

apt-get update
apt-get upgrade -y

# Install common dependencies
apt-get install -y \
    apt-transport-https \
    build-essential \
    ca-certificates \
    curl \
    git \
    gnupg \
    gpg \
    lsb-release \
    nano \
    openssh-client \
    rsync \
    sudo \
    unzip \
    software-properties-common \
   
# Set the locale
apt-get install -y locales

sed -i -e 's/# en_US.UTF-8 UTF-8/en_US.UTF-8 UTF-8/' /etc/locale.gen
locale-gen

>> /etc/environment cat <<EOF
LANG="en_US.UTF-8"
LANGUAGE="en_US:en"
LC_ALL="en_US.UTF-8"
EOF

apt-get autoremove -y
apt-get clean autoclean
rm -rf /var/lib/{apt,dpkg,cache,log}