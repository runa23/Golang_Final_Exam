// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: wallet/proto/wallet_service/v1/wallet.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	WalletService_GetWallets_FullMethodName        = "/proto.user_service.v1.WalletService/GetWallets"
	WalletService_GetWalletByID_FullMethodName     = "/proto.user_service.v1.WalletService/GetWalletByID"
	WalletService_CreateWallet_FullMethodName      = "/proto.user_service.v1.WalletService/CreateWallet"
	WalletService_CreateTransaction_FullMethodName = "/proto.user_service.v1.WalletService/CreateTransaction"
	WalletService_TopUpWallet_FullMethodName       = "/proto.user_service.v1.WalletService/TopUpWallet"
	WalletService_TransferWallet_FullMethodName    = "/proto.user_service.v1.WalletService/TransferWallet"
	WalletService_GetWalletBalance_FullMethodName  = "/proto.user_service.v1.WalletService/GetWalletBalance"
	WalletService_GetTransactions_FullMethodName   = "/proto.user_service.v1.WalletService/GetTransactions"
)

// WalletServiceClient is the client API for WalletService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletServiceClient interface {
	GetWallets(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetWalletsResponse, error)
	GetWalletByID(ctx context.Context, in *GetWalletByIDRequest, opts ...grpc.CallOption) (*GetWalletByIDResponse, error)
	CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*WalletMutationResponse, error)
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*WalletMutationResponse, error)
	TopUpWallet(ctx context.Context, in *TopUpWalletRequest, opts ...grpc.CallOption) (*WalletMutationResponse, error)
	TransferWallet(ctx context.Context, in *TransferWalletRequest, opts ...grpc.CallOption) (*WalletMutationResponse, error)
	GetWalletBalance(ctx context.Context, in *GetWalletBalanceRequest, opts ...grpc.CallOption) (*GetWalletBalanceResponse, error)
	GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error)
}

type walletServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletServiceClient(cc grpc.ClientConnInterface) WalletServiceClient {
	return &walletServiceClient{cc}
}

func (c *walletServiceClient) GetWallets(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetWalletsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWalletsResponse)
	err := c.cc.Invoke(ctx, WalletService_GetWallets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) GetWalletByID(ctx context.Context, in *GetWalletByIDRequest, opts ...grpc.CallOption) (*GetWalletByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWalletByIDResponse)
	err := c.cc.Invoke(ctx, WalletService_GetWalletByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*WalletMutationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WalletMutationResponse)
	err := c.cc.Invoke(ctx, WalletService_CreateWallet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*WalletMutationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WalletMutationResponse)
	err := c.cc.Invoke(ctx, WalletService_CreateTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) TopUpWallet(ctx context.Context, in *TopUpWalletRequest, opts ...grpc.CallOption) (*WalletMutationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WalletMutationResponse)
	err := c.cc.Invoke(ctx, WalletService_TopUpWallet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) TransferWallet(ctx context.Context, in *TransferWalletRequest, opts ...grpc.CallOption) (*WalletMutationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WalletMutationResponse)
	err := c.cc.Invoke(ctx, WalletService_TransferWallet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) GetWalletBalance(ctx context.Context, in *GetWalletBalanceRequest, opts ...grpc.CallOption) (*GetWalletBalanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWalletBalanceResponse)
	err := c.cc.Invoke(ctx, WalletService_GetWalletBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTransactionsResponse)
	err := c.cc.Invoke(ctx, WalletService_GetTransactions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServiceServer is the server API for WalletService service.
// All implementations must embed UnimplementedWalletServiceServer
// for forward compatibility.
type WalletServiceServer interface {
	GetWallets(context.Context, *emptypb.Empty) (*GetWalletsResponse, error)
	GetWalletByID(context.Context, *GetWalletByIDRequest) (*GetWalletByIDResponse, error)
	CreateWallet(context.Context, *CreateWalletRequest) (*WalletMutationResponse, error)
	CreateTransaction(context.Context, *CreateTransactionRequest) (*WalletMutationResponse, error)
	TopUpWallet(context.Context, *TopUpWalletRequest) (*WalletMutationResponse, error)
	TransferWallet(context.Context, *TransferWalletRequest) (*WalletMutationResponse, error)
	GetWalletBalance(context.Context, *GetWalletBalanceRequest) (*GetWalletBalanceResponse, error)
	GetTransactions(context.Context, *GetTransactionsRequest) (*GetTransactionsResponse, error)
	mustEmbedUnimplementedWalletServiceServer()
}

// UnimplementedWalletServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWalletServiceServer struct{}

func (UnimplementedWalletServiceServer) GetWallets(context.Context, *emptypb.Empty) (*GetWalletsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWallets not implemented")
}
func (UnimplementedWalletServiceServer) GetWalletByID(context.Context, *GetWalletByIDRequest) (*GetWalletByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWalletByID not implemented")
}
func (UnimplementedWalletServiceServer) CreateWallet(context.Context, *CreateWalletRequest) (*WalletMutationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWallet not implemented")
}
func (UnimplementedWalletServiceServer) CreateTransaction(context.Context, *CreateTransactionRequest) (*WalletMutationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}
func (UnimplementedWalletServiceServer) TopUpWallet(context.Context, *TopUpWalletRequest) (*WalletMutationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopUpWallet not implemented")
}
func (UnimplementedWalletServiceServer) TransferWallet(context.Context, *TransferWalletRequest) (*WalletMutationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferWallet not implemented")
}
func (UnimplementedWalletServiceServer) GetWalletBalance(context.Context, *GetWalletBalanceRequest) (*GetWalletBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWalletBalance not implemented")
}
func (UnimplementedWalletServiceServer) GetTransactions(context.Context, *GetTransactionsRequest) (*GetTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactions not implemented")
}
func (UnimplementedWalletServiceServer) mustEmbedUnimplementedWalletServiceServer() {}
func (UnimplementedWalletServiceServer) testEmbeddedByValue()                       {}

// UnsafeWalletServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletServiceServer will
// result in compilation errors.
type UnsafeWalletServiceServer interface {
	mustEmbedUnimplementedWalletServiceServer()
}

func RegisterWalletServiceServer(s grpc.ServiceRegistrar, srv WalletServiceServer) {
	// If the following call pancis, it indicates UnimplementedWalletServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WalletService_ServiceDesc, srv)
}

func _WalletService_GetWallets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetWallets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletService_GetWallets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetWallets(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_GetWalletByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWalletByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetWalletByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletService_GetWalletByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetWalletByID(ctx, req.(*GetWalletByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_CreateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).CreateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletService_CreateWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).CreateWallet(ctx, req.(*CreateWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletService_CreateTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_TopUpWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopUpWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).TopUpWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletService_TopUpWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).TopUpWallet(ctx, req.(*TopUpWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_TransferWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).TransferWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletService_TransferWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).TransferWallet(ctx, req.(*TransferWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_GetWalletBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWalletBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetWalletBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletService_GetWalletBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetWalletBalance(ctx, req.(*GetWalletBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_GetTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletService_GetTransactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetTransactions(ctx, req.(*GetTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WalletService_ServiceDesc is the grpc.ServiceDesc for WalletService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WalletService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.user_service.v1.WalletService",
	HandlerType: (*WalletServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWallets",
			Handler:    _WalletService_GetWallets_Handler,
		},
		{
			MethodName: "GetWalletByID",
			Handler:    _WalletService_GetWalletByID_Handler,
		},
		{
			MethodName: "CreateWallet",
			Handler:    _WalletService_CreateWallet_Handler,
		},
		{
			MethodName: "CreateTransaction",
			Handler:    _WalletService_CreateTransaction_Handler,
		},
		{
			MethodName: "TopUpWallet",
			Handler:    _WalletService_TopUpWallet_Handler,
		},
		{
			MethodName: "TransferWallet",
			Handler:    _WalletService_TransferWallet_Handler,
		},
		{
			MethodName: "GetWalletBalance",
			Handler:    _WalletService_GetWalletBalance_Handler,
		},
		{
			MethodName: "GetTransactions",
			Handler:    _WalletService_GetTransactions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet/proto/wallet_service/v1/wallet.proto",
}
