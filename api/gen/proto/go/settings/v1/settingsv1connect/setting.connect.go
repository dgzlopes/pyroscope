// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: settings/v1/setting.proto

package settingsv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/grafana/pyroscope/api/gen/proto/go/settings/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// SettingsServiceName is the fully-qualified name of the SettingsService service.
	SettingsServiceName = "settings.v1.SettingsService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SettingsServiceGetProcedure is the fully-qualified name of the SettingsService's Get RPC.
	SettingsServiceGetProcedure = "/settings.v1.SettingsService/Get"
	// SettingsServiceSetProcedure is the fully-qualified name of the SettingsService's Set RPC.
	SettingsServiceSetProcedure = "/settings.v1.SettingsService/Set"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	settingsServiceServiceDescriptor   = v1.File_settings_v1_setting_proto.Services().ByName("SettingsService")
	settingsServiceGetMethodDescriptor = settingsServiceServiceDescriptor.Methods().ByName("Get")
	settingsServiceSetMethodDescriptor = settingsServiceServiceDescriptor.Methods().ByName("Set")
)

// SettingsServiceClient is a client for the settings.v1.SettingsService service.
type SettingsServiceClient interface {
	Get(context.Context, *connect.Request[v1.GetSettingsRequest]) (*connect.Response[v1.GetSettingsResponse], error)
	Set(context.Context, *connect.Request[v1.SetSettingsRequest]) (*connect.Response[v1.SetSettingsResponse], error)
}

// NewSettingsServiceClient constructs a client for the settings.v1.SettingsService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSettingsServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) SettingsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &settingsServiceClient{
		get: connect.NewClient[v1.GetSettingsRequest, v1.GetSettingsResponse](
			httpClient,
			baseURL+SettingsServiceGetProcedure,
			connect.WithSchema(settingsServiceGetMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		set: connect.NewClient[v1.SetSettingsRequest, v1.SetSettingsResponse](
			httpClient,
			baseURL+SettingsServiceSetProcedure,
			connect.WithSchema(settingsServiceSetMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// settingsServiceClient implements SettingsServiceClient.
type settingsServiceClient struct {
	get *connect.Client[v1.GetSettingsRequest, v1.GetSettingsResponse]
	set *connect.Client[v1.SetSettingsRequest, v1.SetSettingsResponse]
}

// Get calls settings.v1.SettingsService.Get.
func (c *settingsServiceClient) Get(ctx context.Context, req *connect.Request[v1.GetSettingsRequest]) (*connect.Response[v1.GetSettingsResponse], error) {
	return c.get.CallUnary(ctx, req)
}

// Set calls settings.v1.SettingsService.Set.
func (c *settingsServiceClient) Set(ctx context.Context, req *connect.Request[v1.SetSettingsRequest]) (*connect.Response[v1.SetSettingsResponse], error) {
	return c.set.CallUnary(ctx, req)
}

// SettingsServiceHandler is an implementation of the settings.v1.SettingsService service.
type SettingsServiceHandler interface {
	Get(context.Context, *connect.Request[v1.GetSettingsRequest]) (*connect.Response[v1.GetSettingsResponse], error)
	Set(context.Context, *connect.Request[v1.SetSettingsRequest]) (*connect.Response[v1.SetSettingsResponse], error)
}

// NewSettingsServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSettingsServiceHandler(svc SettingsServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	settingsServiceGetHandler := connect.NewUnaryHandler(
		SettingsServiceGetProcedure,
		svc.Get,
		connect.WithSchema(settingsServiceGetMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	settingsServiceSetHandler := connect.NewUnaryHandler(
		SettingsServiceSetProcedure,
		svc.Set,
		connect.WithSchema(settingsServiceSetMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/settings.v1.SettingsService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SettingsServiceGetProcedure:
			settingsServiceGetHandler.ServeHTTP(w, r)
		case SettingsServiceSetProcedure:
			settingsServiceSetHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSettingsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSettingsServiceHandler struct{}

func (UnimplementedSettingsServiceHandler) Get(context.Context, *connect.Request[v1.GetSettingsRequest]) (*connect.Response[v1.GetSettingsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("settings.v1.SettingsService.Get is not implemented"))
}

func (UnimplementedSettingsServiceHandler) Set(context.Context, *connect.Request[v1.SetSettingsRequest]) (*connect.Response[v1.SetSettingsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("settings.v1.SettingsService.Set is not implemented"))
}
