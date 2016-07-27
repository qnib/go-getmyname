#!/bin/bash
set -ex

if [ -f /etc/os-release ];then
    . /etc/os-release
    if [ "X${ID}" != "Xalpine" ];then
      ID=Linux
    fi
else
    ID=$(uname -s)
fi
go get -d
go build -o $(basename $(pwd))_$(git describe --abbrev=0 --tags)_${ID}
