// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.23.3
// source: baby/v1/baby.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationBabyExample = "/api.baby.v1.Baby/Example"
const OperationBabyGetStoryList = "/api.baby.v1.Baby/GetStoryList"
const OperationBabyGetUser = "/api.baby.v1.Baby/GetUser"

type BabyHTTPServer interface {
	Example(context.Context, *Null) (*emptypb.Empty, error)
	GetStoryList(context.Context, *GetStoryListRequest) (*GetStoryListReply, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
}

func RegisterBabyHTTPServer(s *http.Server, srv BabyHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/baby/example", _Baby_Example0_HTTP_Handler(srv))
	r.GET("/v1/baby/getuser", _Baby_GetUser1_HTTP_Handler(srv))
	r.GET("/v1/baby/storylist", _Baby_GetStoryList0_HTTP_Handler(srv))
}

func _Baby_Example0_HTTP_Handler(srv BabyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in Null
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBabyExample)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Example(ctx, req.(*Null))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Baby_GetUser1_HTTP_Handler(srv BabyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBabyGetUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserReply)
		return ctx.Result(200, reply)
	}
}

func _Baby_GetStoryList0_HTTP_Handler(srv BabyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetStoryListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBabyGetStoryList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetStoryList(ctx, req.(*GetStoryListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetStoryListReply)
		return ctx.Result(200, reply)
	}
}

type BabyHTTPClient interface {
	Example(ctx context.Context, req *Null, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetStoryList(ctx context.Context, req *GetStoryListRequest, opts ...http.CallOption) (rsp *GetStoryListReply, err error)
	GetUser(ctx context.Context, req *GetUserRequest, opts ...http.CallOption) (rsp *GetUserReply, err error)
}

type BabyHTTPClientImpl struct {
	cc *http.Client
}

func NewBabyHTTPClient(client *http.Client) BabyHTTPClient {
	return &BabyHTTPClientImpl{client}
}

func (c *BabyHTTPClientImpl) Example(ctx context.Context, in *Null, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/v1/baby/example"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBabyExample))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BabyHTTPClientImpl) GetStoryList(ctx context.Context, in *GetStoryListRequest, opts ...http.CallOption) (*GetStoryListReply, error) {
	var out GetStoryListReply
	pattern := "/v1/baby/storylist"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBabyGetStoryList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BabyHTTPClientImpl) GetUser(ctx context.Context, in *GetUserRequest, opts ...http.CallOption) (*GetUserReply, error) {
	var out GetUserReply
	pattern := "/v1/baby/getuser"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBabyGetUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
