FROM mcr.microsoft.com/devcontainers/base:ubuntu-22.04

RUN export DEBIAN_FRONTEND=noninteractive \
    && apt-get update \
    && apt-get -y install --no-install-recommends \
        git \
        python3 \
        python3-pip \
        wget \
        curl \
    && apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*

RUN pip3 install --upgrade pip \
    && pip3 install ansible ansible-lint