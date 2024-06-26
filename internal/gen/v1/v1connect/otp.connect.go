// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: v1/otp.proto

package v1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	v1 "otp-service/internal/gen/v1"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// OTPServiceName is the fully-qualified name of the OTPService service.
	OTPServiceName = "com.service.otp.OTPService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// OTPServiceGenerateOTPProcedure is the fully-qualified name of the OTPService's generateOTP RPC.
	OTPServiceGenerateOTPProcedure = "/com.service.otp.OTPService/generateOTP"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	oTPServiceServiceDescriptor           = v1.File_v1_otp_proto.Services().ByName("OTPService")
	oTPServiceGenerateOTPMethodDescriptor = oTPServiceServiceDescriptor.Methods().ByName("generateOTP")
)

// OTPServiceClient is a client for the com.service.otp.OTPService service.
type OTPServiceClient interface {
	GenerateOTP(context.Context, *connect.Request[v1.GenerateOTPRequest]) (*connect.Response[v1.GenerateOTPResponse], error)
}

// NewOTPServiceClient constructs a client for the com.service.otp.OTPService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewOTPServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) OTPServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &oTPServiceClient{
		generateOTP: connect.NewClient[v1.GenerateOTPRequest, v1.GenerateOTPResponse](
			httpClient,
			baseURL+OTPServiceGenerateOTPProcedure,
			connect.WithSchema(oTPServiceGenerateOTPMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// oTPServiceClient implements OTPServiceClient.
type oTPServiceClient struct {
	generateOTP *connect.Client[v1.GenerateOTPRequest, v1.GenerateOTPResponse]
}

// GenerateOTP calls com.service.otp.OTPService.generateOTP.
func (c *oTPServiceClient) GenerateOTP(ctx context.Context, req *connect.Request[v1.GenerateOTPRequest]) (*connect.Response[v1.GenerateOTPResponse], error) {
	return c.generateOTP.CallUnary(ctx, req)
}

// OTPServiceHandler is an implementation of the com.service.otp.OTPService service.
type OTPServiceHandler interface {
	GenerateOTP(context.Context, *connect.Request[v1.GenerateOTPRequest]) (*connect.Response[v1.GenerateOTPResponse], error)
}

// NewOTPServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewOTPServiceHandler(svc OTPServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	oTPServiceGenerateOTPHandler := connect.NewUnaryHandler(
		OTPServiceGenerateOTPProcedure,
		svc.GenerateOTP,
		connect.WithSchema(oTPServiceGenerateOTPMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/com.service.otp.OTPService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case OTPServiceGenerateOTPProcedure:
			oTPServiceGenerateOTPHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedOTPServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedOTPServiceHandler struct{}

func (UnimplementedOTPServiceHandler) GenerateOTP(context.Context, *connect.Request[v1.GenerateOTPRequest]) (*connect.Response[v1.GenerateOTPResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("com.service.otp.OTPService.generateOTP is not implemented"))
}
