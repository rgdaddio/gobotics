CC      = /home/rdaddio/myGCC_run/myGCC_out/bin/gccgo
CFLAGS  = -g
RM      = rm -f
BUILDER := /usr/local/opt/go@1.13/bin/go
#export GOPATH := $(shell pwd)/lib
#export GOROOT := $(shell pwd)/go

default: all

all: gobotics

build_server:
	$(BUILDER)  build -o gobotics_server -work -x server/*

gobotics: gobotics.go
#	if test ! -s go1.10.1.linux-amd64.tar.gz ;\
#	then \
#		rm -f go1.10.1.linux-amd64.tar* ; \
#		wget https://storage.googleapis.com/golang/go1.10.1.linux-amd64.tar.gz ; \
#		tar -xvzf go1.10.1.linux-amd64.tar.gz ; \
#	fi;

	$(BUILDER) get -u github.com/mattn/go-sqlite3

	$(BUILDER) get -u -d gobot.io/x/gobot/...

	$(BUILDER)  build -work -x gobotics.go client.go
	$(BUILDER)  build -o gobotics_server -work -x server/* 

clean veryclean:
	$(RM) gobotics
	$(RM) gobotics_server
