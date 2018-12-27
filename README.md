# Metadium Proxy

[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/metadium/go-delegator/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/metadium/go-delegator)](https://goreportcard.com/report/github.com/metadium/go-delegator)
[![GoDoc](https://godoc.org/github.com/metadium/go-delegator?status.svg)](https://godoc.org/github.com/metadium/go-delegator)

> Main functions are MetaID CRUD and relaying JSON-RPC result with metadium node. In addition, this project try porting web3 to Golang. Furthermore it applied IPFS here to overwhelm limited storage of blockchain.

## Contents
* [Features](#markdown-header-features)
* [Prerequisite](#markdown-header-prerequisite)
* [Build](#markdown-header-build)
* [Test](#markdown-header-test)
* [Usage](#markdown-header-usage)
* [Deploy (for AWS Lambda)](#markdown-header-deploy-for-aws-lambda)
* [Documentation](#markdown-header-documentation)
* [Reference](#markdown-header-reference)
* [License](#markdown-header-license)

## Features
1. JSON-RPC relay with metadium node
2. Proofs for sign and merkle tree such as Ecrecover, DeriveSha, VerifyProof and so on
3. Sign, SignTx with encrypted private key on DynamoDB
4. IPFS
5. fromWei, toWei written in Golang

## Prerequisite
0. Go
    - Install at https://golang.org/doc/install
1. Dep
    - Install
    ```
    brew install dep
    ```
2. Docker
    - Install at https://docs.docker.com/install
3. xgo
    - because of compilation for C code used in go-ethereum, we need improved cross-compiler
    
```shell
docker pull karalabe/xgo-latest
go get github.com/karalabe/xgo
```

## Build
1. Move to root directory of this repo
2. Build on your preference

  - In case of Lambda that is cross-compile,
`make` or `make lambda`

  - In case of Lambda with remote branch,
`make branch=master remote`

  - In case of compile for local machine,
`make local`

## Test
1. Move each module directory such as json, rpc and so on
2. Run testunit
    ```
    go test -v
    ```
    
## Usage
1. $> proxy [KEY_JSON_PATH] -log_lev=debug -log_out=/log/proxy.log -log_fmt=json
2. $> proxy [KEY_JSON_PATH] [KEY_JSON_PASSPHRASE] -log_lev=debug -log_out=/log/proxy.log -log_fmt=json
    - ```log_lev```, ```log_out```, ```log_fmt```, ```log_bot_token``` and ```log_bot_chatid``` are optional
    - description:
        * log_lev: log level
        * log_out: log output location
        * log_fmt: log format, text or JSON
        * log_bot_token: telegram access token
        * log_bot_chatid: telegram chat ID
    - default:
        * log_lev: info
        * log_out: stdout
        * log_fmt: text


## Deploy (for AWS Lambda)
1. Set Lambda on AWS
    - Function package: compressed binary file in $GOPATH/src/{repo}/bin
    - Handler: proxy (binary file name, it is optional)
    - Runtime: Go 1.x
    - (Optional) Include DynamoDB execution role to Lambda execution role
2. Set API Gateway as proxy on AWS
3. Add API Gateway as Lambda trigger
4. Add CloudWatch Logs

## Documentation
1. Execute godoc -http like below
    ```
    godoc -http=:6060
    ```  
2. Open url
    - http://localhost:6060/pkg/github.com/metadium/go-delegator/
    - If you change port at 1., it should be applied to url

## Reference
[1] AWS Lambda Go, https://github.com/aws/aws-lambda-go

[2] Go ethereum, https://github.com/ethereum/go-ethereum

[3] IPFS, https://ipfs.io/

[4] IPFS API, https://github.com/ipfs/go-ipfs-api

[5] Cross compiling for ethereum, https://github.com/ethereum/go-ethereum/wiki/Cross-compiling-Ethereum

[6] Xgo, https://github.com/karalabe/xgo

[7] Dep, https://github.com/golang/dep

## License
MIT
