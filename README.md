# Metadium Proxy on AWS Lambda
Main functions are MetaID registration/update/recover and relaying JSON-RPC result with metadium node.

In addition, this project try porting web3 to Golang.

Furthermore it applied IPFS here to overwhelm limited storage of blockchain.

# Features
1. JSON-RPC relay with metadium node
2. Proofs for sign and merkle tree such as Ecrecover, DeriveSha, VerifyProof and so on
3. Sign, SignTx with encrypted private key on DynamoDB
4. IPFS
5. fromWei, toWei written in Golang

# Prerequisite
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
    - because of C compile in go-ethereum, we need improved cross-compiler
    
```shell
docker pull karalabe/xgo-latest
go get github.com/karalabe/xgo
```

# Build
```shell
cd $GOPATH/src/{repo start with bitbucket.org}
make
```

# Test
1. Move each module directory such as json, rpc and so on
2. Run testunit
    ```
    go test -v
    ```

# Deploy
1. Set Lambda on AWS
    - Function package: compressed binary file in $GOPATH/src/{repo}/bin
    - Handler: proxy (binary file name, it is optional)
    - Runtime: Go 1.x
    - (Optional) Include DynamoDB execution role to Lambda execution role
2. Set API Gateway as proxy on AWS
3. Add API Gateway as Lambda trigger
4. Add CloudWatch Logs

# Documentation
1. Execute godoc -http like below
    ```
    godoc -http=:6060
    ```  
2. Open url
    - http://localhost:6060/pkg/bitbucket.org/coinplugin/proxy/
    - If you change port at 1., it should be applied to url

# Reference
[1] AWS Lambda Go, https://github.com/aws/aws-lambda-go

[2] Go ethereum, https://github.com/ethereum/go-ethereum

[3] IPFS, https://ipfs.io/

[4] IPFS API, https://github.com/ipfs/go-ipfs-api

[5] Cross compliling for ethereum, https://github.com/ethereum/go-ethereum/wiki/Cross-compiling-Ethereum

[6] Xgo, https://github.com/karalabe/xgo

[7] Dep, https://github.com/golang/dep

# License
MIT
