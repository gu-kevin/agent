package api

import (
	"context"
	"io/ioutil"
	"reflect"

	"github.com/gogo/protobuf/proto"
	"go.uber.org/fx"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/encoding/protobuf"
	"go.uber.org/yarpc/encoding/protobuf/reflection"
)

var _ = ioutil.NopCloser

// GetBuildkiteAPIYARPCClient is the YARPC client-side interface for the GetBuildkiteAPI service.
type GetBuildkiteAPIYARPCClient interface {
	Ping(context.Context, *PingRequest, ...yarpc.CallOption) (*PingResponse, error)
}

// NewGetBuildkiteAPIYARPCClient builds a new YARPC client for the GetBuildkiteAPI service.
func NewGetBuildkiteAPIYARPCClient(clientConfig transport.ClientConfig, options ...protobuf.ClientOption) GetBuildkiteAPIYARPCClient {
	return &_GetBuildkiteAPIYARPCCaller{protobuf.NewStreamClient(
		protobuf.ClientParams{
			ServiceName:  "uber.proto.atg.buildkite.v1.GetBuildkiteAPI",
			ClientConfig: clientConfig,
			Options:      options,
		},
	)}
}

// GetBuildkiteAPIYARPCServer is the YARPC server-side interface for the GetBuildkiteAPI service.
type GetBuildkiteAPIYARPCServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
}

// BuildGetBuildkiteAPIYARPCProcedures prepares an implementation of the GetBuildkiteAPI service for YARPC registration.
func BuildGetBuildkiteAPIYARPCProcedures(server GetBuildkiteAPIYARPCServer) []transport.Procedure {
	handler := &_GetBuildkiteAPIYARPCHandler{server}
	return protobuf.BuildProcedures(
		protobuf.BuildProceduresParams{
			ServiceName: "uber.proto.atg.buildkite.v1.GetBuildkiteAPI",
			UnaryHandlerParams: []protobuf.BuildProceduresUnaryHandlerParams{
				{
					MethodName: "Ping",
					Handler: protobuf.NewUnaryHandler(
						protobuf.UnaryHandlerParams{
							Handle:     handler.Ping,
							NewRequest: newGetBuildkiteAPIServicePingYARPCRequest,
						},
					),
				},
			},
			OnewayHandlerParams: []protobuf.BuildProceduresOnewayHandlerParams{},
			StreamHandlerParams: []protobuf.BuildProceduresStreamHandlerParams{},
		},
	)
}

// FxGetBuildkiteAPIYARPCClientParams defines the input
// for NewFxGetBuildkiteAPIYARPCClient. It provides the
// paramaters to get a GetBuildkiteAPIYARPCClient in an
// Fx application.
type FxGetBuildkiteAPIYARPCClientParams struct {
	fx.In

	Provider yarpc.ClientConfig
}

// FxGetBuildkiteAPIYARPCClientResult defines the output
// of NewFxGetBuildkiteAPIYARPCClient. It provides a
// GetBuildkiteAPIYARPCClient to an Fx application.
type FxGetBuildkiteAPIYARPCClientResult struct {
	fx.Out

	Client GetBuildkiteAPIYARPCClient

	// We are using an fx.Out struct here instead of just returning a client
	// so that we can add more values or add named versions of the client in
	// the future without breaking any existing code.
}

// NewFxGetBuildkiteAPIYARPCClient provides a GetBuildkiteAPIYARPCClient
// to an Fx application using the given name for routing.
//
//  fx.Provide(
//    buildkitev1.NewFxGetBuildkiteAPIYARPCClient("service-name"),
//    ...
//  )
func NewFxGetBuildkiteAPIYARPCClient(name string, options ...protobuf.ClientOption) interface{} {
	return func(params FxGetBuildkiteAPIYARPCClientParams) FxGetBuildkiteAPIYARPCClientResult {
		return FxGetBuildkiteAPIYARPCClientResult{
			Client: NewGetBuildkiteAPIYARPCClient(params.Provider.ClientConfig(name), options...),
		}
	}
}

// FxGetBuildkiteAPIYARPCProceduresParams defines the input
// for NewFxGetBuildkiteAPIYARPCProcedures. It provides the
// paramaters to get GetBuildkiteAPIYARPCServer procedures in an
// Fx application.
type FxGetBuildkiteAPIYARPCProceduresParams struct {
	fx.In

	Server GetBuildkiteAPIYARPCServer
}

// FxGetBuildkiteAPIYARPCProceduresResult defines the output
// of NewFxGetBuildkiteAPIYARPCProcedures. It provides
// GetBuildkiteAPIYARPCServer procedures to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group.
// Dig 1.2 or newer must be used for this feature to work.
type FxGetBuildkiteAPIYARPCProceduresResult struct {
	fx.Out

	Procedures     []transport.Procedure `group:"yarpcfx"`
	ReflectionMeta reflection.ServerMeta `group:"yarpcfx"`
}

// NewFxGetBuildkiteAPIYARPCProcedures provides GetBuildkiteAPIYARPCServer procedures to an Fx application.
// It expects a GetBuildkiteAPIYARPCServer to be present in the container.
//
//  fx.Provide(
//    buildkitev1.NewFxGetBuildkiteAPIYARPCProcedures(),
//    ...
//  )
func NewFxGetBuildkiteAPIYARPCProcedures() interface{} {
	return func(params FxGetBuildkiteAPIYARPCProceduresParams) FxGetBuildkiteAPIYARPCProceduresResult {
		return FxGetBuildkiteAPIYARPCProceduresResult{
			Procedures: BuildGetBuildkiteAPIYARPCProcedures(params.Server),
			ReflectionMeta: reflection.ServerMeta{
				ServiceName:     "uber.proto.atg.buildkite.v1.GetBuildkiteAPI",
				FileDescriptors: yarpcFileDescriptorClosure7672be1ea7f0e766,
			},
		}
	}
}

type _GetBuildkiteAPIYARPCCaller struct {
	streamClient protobuf.StreamClient
}

func (c *_GetBuildkiteAPIYARPCCaller) Ping(ctx context.Context, request *PingRequest, options ...yarpc.CallOption) (*PingResponse, error) {
	responseMessage, err := c.streamClient.Call(ctx, "Ping", request, newGetBuildkiteAPIServicePingYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*PingResponse)
	if !ok {
		return nil, protobuf.CastError(emptyGetBuildkiteAPIServicePingYARPCResponse, responseMessage)
	}
	return response, err
}

type _GetBuildkiteAPIYARPCHandler struct {
	server GetBuildkiteAPIYARPCServer
}

func (h *_GetBuildkiteAPIYARPCHandler) Ping(ctx context.Context, requestMessage proto.Message) (proto.Message, error) {
	var request *PingRequest
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*PingRequest)
		if !ok {
			return nil, protobuf.CastError(emptyGetBuildkiteAPIServicePingYARPCRequest, requestMessage)
		}
	}
	response, err := h.server.Ping(ctx, request)
	if response == nil {
		return nil, err
	}
	return response, err
}

func newGetBuildkiteAPIServicePingYARPCRequest() proto.Message {
	return &PingRequest{}
}

func newGetBuildkiteAPIServicePingYARPCResponse() proto.Message {
	return &PingResponse{}
}

var (
	emptyGetBuildkiteAPIServicePingYARPCRequest  = &PingRequest{}
	emptyGetBuildkiteAPIServicePingYARPCResponse = &PingResponse{}
)

var yarpcFileDescriptorClosure7672be1ea7f0e766 = [][]byte{
	// uber/proto/atg/buildkite/v1/get_buildkite_api.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8f, 0xb1, 0x4e, 0xc3, 0x30,
		0x10, 0x86, 0x95, 0xaa, 0x42, 0xe0, 0x20, 0x21, 0x45, 0x0c, 0x28, 0x1d, 0x8a, 0x3a, 0x15, 0x86,
		0xb3, 0xd2, 0xce, 0x0c, 0xc9, 0x82, 0xd8, 0xac, 0x88, 0x76, 0x80, 0x48, 0x55, 0x02, 0xc6, 0xb2,
		0x68, 0x62, 0x93, 0x5c, 0xf2, 0x0e, 0xbc, 0x06, 0x23, 0x8f, 0xc2, 0x53, 0x21, 0xdb, 0xc4, 0x82,
		0xa5, 0x62, 0xbb, 0xe4, 0xbe, 0xef, 0xbf, 0xdf, 0x64, 0xdd, 0x57, 0xbc, 0xa5, 0xba, 0x55, 0xa8,
		0x68, 0x89, 0x82, 0x56, 0xbd, 0xdc, 0x3f, 0xbf, 0x4a, 0xe4, 0x74, 0x48, 0xa8, 0xe0, 0xb8, 0xf3,
		0x3f, 0x76, 0xa5, 0x96, 0x60, 0xc1, 0x68, 0x66, 0x24, 0x37, 0x43, 0x89, 0x02, 0x3c, 0x03, 0x43,
		0x12, 0xcf, 0x85, 0x52, 0x62, 0xcf, 0x5d, 0x66, 0xd5, 0xbf, 0x50, 0x94, 0x35, 0xef, 0xb0, 0xac,
		0xb5, 0x33, 0x16, 0x37, 0x24, 0x64, 0xb2, 0x11, 0x39, 0x7f, 0xeb, 0x79, 0x87, 0x11, 0x90, 0xa9,
		0x21, 0x2e, 0x82, 0xcb, 0x60, 0x19, 0xae, 0x62, 0x70, 0x3a, 0x8c, 0x3a, 0xdc, 0x8f, 0x7a, 0x6e,
		0xb9, 0xc5, 0x35, 0x39, 0x75, 0x7a, 0xa7, 0x55, 0xd3, 0xf1, 0x28, 0x26, 0xc7, 0xed, 0xcf, 0x6c,
		0x33, 0x4e, 0x72, 0xff, 0xbd, 0x6a, 0xc8, 0xd9, 0x2d, 0xc7, 0x6c, 0xac, 0x97, 0xb2, 0xbb, 0xe8,
		0x91, 0x4c, 0x8d, 0x1e, 0x2d, 0xe1, 0xc0, 0x23, 0xe0, 0x57, 0xc1, 0xf8, 0xea, 0x1f, 0xa4, 0xbb,
		0x97, 0xbd, 0x07, 0x64, 0xfe, 0xa4, 0xea, 0x43, 0x42, 0x76, 0xfe, 0xa7, 0x91, 0x96, 0xcc, 0x60,
		0x2c, 0x78, 0x08, 0x3d, 0x35, 0x24, 0x1f, 0x93, 0xe9, 0x86, 0xa5, 0xd9, 0xe7, 0x64, 0xb6, 0x31,
		0x49, 0x16, 0x81, 0x14, 0x05, 0x78, 0x11, 0xb6, 0xc9, 0x97, 0xdb, 0x16, 0x76, 0x5b, 0xa4, 0x28,
		0x0a, 0xbf, 0x2d, 0xb6, 0x49, 0x75, 0x64, 0xef, 0xaf, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x77,
		0xf2, 0x72, 0xa9, 0xe2, 0x01, 0x00, 0x00,
	},
	// google/protobuf/timestamp.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0xcf, 0xcf, 0x4f,
		0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0xc9, 0xcc, 0x4d,
		0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0xd0, 0x03, 0x0b, 0x09, 0xf1, 0x43, 0x14, 0xe8, 0xc1, 0x14, 0x28,
		0x59, 0x73, 0x71, 0x86, 0xc0, 0xd4, 0x08, 0x49, 0x70, 0xb1, 0x17, 0xa7, 0x26, 0xe7, 0xe7, 0xa5,
		0x14, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0xc1, 0xb8, 0x42, 0x22, 0x5c, 0xac, 0x79, 0x89,
		0x79, 0xf9, 0xc5, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x10, 0x8e, 0x53, 0x1d, 0x97, 0x70,
		0x72, 0x7e, 0xae, 0x1e, 0x9a, 0x99, 0x4e, 0x7c, 0x70, 0x13, 0x03, 0x40, 0x42, 0x01, 0x8c, 0x51,
		0xda, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xe9, 0xf9, 0x39, 0x89,
		0x79, 0xe9, 0x08, 0x27, 0x16, 0x94, 0x54, 0x16, 0xa4, 0x16, 0x23, 0x5c, 0xfa, 0x83, 0x91, 0x71,
		0x11, 0x13, 0xb3, 0x7b, 0x80, 0xd3, 0x2a, 0x26, 0x39, 0x77, 0x88, 0xc9, 0x01, 0x50, 0xb5, 0x7a,
		0xe1, 0xa9, 0x39, 0x39, 0xde, 0x79, 0xf9, 0xe5, 0x79, 0x21, 0x20, 0x3d, 0x49, 0x6c, 0x60, 0x43,
		0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x77, 0x4a, 0x07, 0xf7, 0x00, 0x00, 0x00,
	},
}

func init() {
	yarpc.RegisterClientBuilder(
		func(clientConfig transport.ClientConfig, structField reflect.StructField) GetBuildkiteAPIYARPCClient {
			return NewGetBuildkiteAPIYARPCClient(clientConfig, protobuf.ClientBuilderOptions(clientConfig, structField)...)
		},
	)
}
