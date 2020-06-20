#export GOPATH := $(shell pwd)/lib
#export GOROOT := $(shell pwd)/go

default: all

all: gobotics

build_server:
	go build -o gobotics_server -work -x cmd/server/*

build_client:
	go build -o gobotics_client -work -x cmd/client/*

gobotics: gobotics.go

	$(BUILDER) get -u github.com/mattn/go-sqlite3

	$(BUILDER) get -u -d gobot.io/x/gobot/...

	build_client
	build_server

clean veryclean:
	$(RM) gobotics
	$(RM) gobotics_server
