// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package rpcserviceiface provides an interface to enable mocking the RPC Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package rpcserviceiface

import (
	"github.com/IBM/ibm-cos-sdk-go/aws"
	"github.com/IBM/ibm-cos-sdk-go/aws/request"
	"github.com/IBM/ibm-cos-sdk-go/private/model/api/codegentest/service/rpcservice"
)

// RPCServiceAPI provides an interface to enable mocking the
// rpcservice.RPCService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // RPC Service.
//    func myFunc(svc rpcserviceiface.RPCServiceAPI) bool {
//        // Make svc.EmptyStream request
//    }
//
// IBM COS SDK Code -- START
//    func main() {
//        sess := session.Must(session.NewSession())
//        svc := rpcservice.New(sess)
//
//        myFunc(svc)
//    }
// IBM COS SDK Code -- END
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockRPCServiceClient struct {
//        rpcserviceiface.RPCServiceAPI
//    }
//    func (m *mockRPCServiceClient) EmptyStream(input *rpcservice.EmptyStreamInput) (*rpcservice.EmptyStreamOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockRPCServiceClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type RPCServiceAPI interface {
	EmptyStream(*rpcservice.EmptyStreamInput) (*rpcservice.EmptyStreamOutput, error)
	EmptyStreamWithContext(aws.Context, *rpcservice.EmptyStreamInput, ...request.Option) (*rpcservice.EmptyStreamOutput, error)
	EmptyStreamRequest(*rpcservice.EmptyStreamInput) (*request.Request, *rpcservice.EmptyStreamOutput)

	GetEventStream(*rpcservice.GetEventStreamInput) (*rpcservice.GetEventStreamOutput, error)
	GetEventStreamWithContext(aws.Context, *rpcservice.GetEventStreamInput, ...request.Option) (*rpcservice.GetEventStreamOutput, error)
	GetEventStreamRequest(*rpcservice.GetEventStreamInput) (*request.Request, *rpcservice.GetEventStreamOutput)

	OtherOperation(*rpcservice.OtherOperationInput) (*rpcservice.OtherOperationOutput, error)
	OtherOperationWithContext(aws.Context, *rpcservice.OtherOperationInput, ...request.Option) (*rpcservice.OtherOperationOutput, error)
	OtherOperationRequest(*rpcservice.OtherOperationInput) (*request.Request, *rpcservice.OtherOperationOutput)
}

var _ RPCServiceAPI = (*rpcservice.RPCService)(nil)
