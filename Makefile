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
			--targets=linux/amd64 -out bin/proxy \
			./
	mv bin/proxy-linux-amd64 bin/proxy

remote:
	xgo --deps=https://gmplib.org/download/gmp/gmp-6.0.0a.tar.bz2 \
			--targets=linux/amd64 -out bin/proxy \
			--branch=$(branch) \
			bitbucket.org/coinplugin/proxy
	mv bin/proxy-linux-amd64 bin/proxy

local:
	go build -o bin/proxy