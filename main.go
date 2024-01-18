package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"go-delegator/config"
	"go-delegator/crypto"
	_ "go-delegator/ipfs"
	"go-delegator/json"
	"go-delegator/log"
	"go-delegator/metaresolver"
	"go-delegator/metaservice"
	"go-delegator/rpc"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/fvbock/endless"
)

const (
	// ParamFuncName is a name indicating function's
	ParamFuncName = "func"
	// Targetnet indicates target network
	Targetnet = rpc.Testnet
)

func handler(req json.RPCRequest) (body string, statusCode int) {
	//log.Info("request:", req.String())
	var resp json.RPCResponse
	var err error
	if metaresolver.Contains(req.Method) {
		// Forward RPC request to metaservice function (v3)
		resp, err = metaresolver.Forward(req)
	} else if metaservice.Contains(req.Method) {
		// Forward RPC request to metaservice function (v2)
		resp, err = metaservice.Forward(req)
	} else {
		// Forward RPC request to Ether node
		var respBody string
		if respBody, err = rpc.GetInstance().DoRPC(req); err == nil {
			// Relay a response from the node
			resp = json.GetRPCResponseFromJSON(respBody)
		}
	}

	statusCode = 200
	if err != nil {
		// In case of server-side RPC fail
		log.Error(err.Error())
		resp.Error = &json.RPCError{
			Code:    -1,
			Message: err.Error(),
		}
		statusCode = 400
	}
	body = resp.String()
	return
}

// lambdaHandler handles APIGatewayProxyRequest as JSON-RPC request
func lambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Validate RPC request
	req := json.GetRPCRequestFromJSON(request.Body)
	if method := request.QueryStringParameters[ParamFuncName]; method != "" {
		req.Method = method
	} else if method := request.PathParameters[ParamFuncName]; method != "" {
		req.Method = method
	}

	respBody, statusCode := handler(req)
	return events.APIGatewayProxyResponse{Body: respBody, StatusCode: statusCode}, nil
}

// httpHandler handles http.Request as JSON-RPC request
func httpHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	log.Info("request:", r.RemoteAddr, string(b))
	req := json.GetRPCRequestFromJSON(string(b))
	respBody, statusCode := handler(req)
	log.Info("response:", r.RemoteAddr, statusCode, respBody)
	w.WriteHeader(statusCode)
	w.Write([]byte(respBody))
}

func help() {
	fmt.Println("USAGE:")
	fmt.Println("  delegator [config_path]")
}

func init() {
	rpc.NetType = Targetnet

	config.ReadConfig()

	// Initialize Crypto with arguments
	go func() {
		crypto.PathChan <- config.Config.AccountPath
		crypto.PassphraseChan <- config.Config.AccountPassword
	}()
	crypto.GetInstance()
}

func main() {
	log.Info("Server starting...")
	if os.Getenv(crypto.IsAwsLambda) != "" {
		log.Info("Ready to start Lambda")
		lambda.Start(lambdaHandler)
	} else {
		log.Info("Ready to start HTTP/HTTPS")
		h := http.NewServeMux()
		h.HandleFunc("/", httpHandler)
		endless.ListenAndServe(":8545", h)
	}
}
