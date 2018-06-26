package predefined

import "fmt"

type Error interface {
	Error() string    // returns the message
	ErrorCode() int32 // returns the code
}

// request is for an unknown service
type methodNotFoundError struct {
	//	service string
	method string
}

func (e *methodNotFoundError) ErrorCode() int32 { return -32601 }

func (e *methodNotFoundError) Error() string {
	//	return fmt.Sprintf("The method %s%s%s does not exist/is not available", e.service, serviceMethodSeparator, e.method)
	return fmt.Sprintf("The method %s does not exist/is not available", e.method)
}

// received message isn't a valid request
type invalidRequestError struct{ message string }

func (e *invalidRequestError) ErrorCode() int32 { return -32600 }

func (e *invalidRequestError) Error() string { return e.message }

// received message is invalid
type invalidMessageError struct{ message string }

func (e *invalidMessageError) ErrorCode() int32 { return -32700 }

func (e *invalidMessageError) Error() string { return e.message }

// unable to decode supplied params, or an invalid number of parameters
type invalidParamsError struct{ message string }

func (e *invalidParamsError) ErrorCode() int32 { return -32602 }

func (e *invalidParamsError) Error() string { return e.message }

type invalidSignatureError struct{ message string }

func (e *invalidSignatureError) ErrorCode() int32 { return -32010 }

func (e *invalidSignatureError) Error() string { return e.message }
