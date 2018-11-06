#!/bin/bash
export GOROOT="/usr/lib/go-1.6"
export GOPATH=$(pwd)
GOROOT=/usr/lib/go-1.6 #gosetup
GOPATH=/home/sebastian/git/go/tdl-go-demo #gosetup
$GOROOT/bin/go build -i -o /tmp/___go_build_main_go $GOPATH/src/main.go #gosetup
/tmp/___go_build_main_go #gosetup
