ifeq ($(GOPATH),)
  export GOPATH=$(HOME)/go
endif
ifeq ($(GOPATH),$(GOROOT))
  export GOPATH=$(GOROOT)/../workspace
endif
export GOBIN=$(GOPATH)/bin

$(shell dep ensure)
$(shell go get github.com/karalabe/xgo)

lambda:
	xgo --deps=https://gmplib.org/download/gmp/gmp-6.0.0a.tar.bz2 \
			--targets=linux/amd64 -out bin/delegator \
			./
	mv bin/delegator-linux-amd64 bin/delegator

remote:
	xgo --deps=https://gmplib.org/download/gmp/gmp-6.0.0a.tar.bz2 \
			--targets=linux/amd64 -out bin/delegator \
			--branch=$(branch) \
			github.com/metadium/go-delegator
	mv bin/delegator-linux-amd64 bin/delegator

local:
	go build -o bin/delegator