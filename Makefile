CC      = gccgo
CFLAGS  = -g
RM      = rm -f
BUILDER := $(CURDIR)/go/bin/go
export GOPATH := $(shell pwd)/lib
export GOROOT := $(shell pwd)/go

default: all

all: gobotics

gobotics: gobotics.go
	if test ! -s go1.5.linux-amd64.tar.gz ;\
	then \
		rm -f go1.5.linux-amd64.tar* ; \
		wget https://storage.googleapis.com/golang/go1.5.linux-amd64.tar.gz ; \
		tar -xvzf go1.5.linux-amd64.tar.gz ; \
	fi;

	$(BUILDER) get golang.org/x/net/context
	$(BUILDER) get -u github.com/mattn/go-sqlite3
	rm -fr /tmp/go-build*
	rm -fr $(shell pwd)/output/*
	$(BUILDER)  build -work -x -compiler gccgo -gccgoflags -pthread github.com/mattn/go-sqlite3/.
	cp -fr /tmp/go-*/github.com/ $(shell pwd)/output/
	cp -fr /tmp/go-*/golang.org/ $(shell pwd)/output/
	$(CC) -v -g -o gobotics gobotics.go -I $(shell pwd)/output $(shell pwd)/output/github.com/mattn/libgo-sqlite3.a $(shell pwd)/output/golang.org/x/net/libcontext.a -pthread -ldl  

clean veryclean:
	$(RM) gobotics
