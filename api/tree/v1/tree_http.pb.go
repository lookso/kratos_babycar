// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.23.3
// source: tree/v1/tree.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationTreeGetTree = "/api.tree.v1.Tree/GetTree"

type TreeHTTPServer interface {
	GetTree(context.Context, *GetTreeRequest) (*GetTreeReply, error)
}

func RegisterTreeHTTPServer(s *http.Server, srv TreeHTTPServer) {
	r := s.Route("/")
	r.GET("/gettree/{id}", _Tree_GetTree0_HTTP_Handler(srv))
}

func _Tree_GetTree0_HTTP_Handler(srv TreeHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTreeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTreeGetTree)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTree(ctx, req.(*GetTreeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTreeReply)
		return ctx.Result(200, reply)
	}
}

type TreeHTTPClient interface {
	GetTree(ctx context.Context, req *GetTreeRequest, opts ...http.CallOption) (rsp *GetTreeReply, err error)
}

type TreeHTTPClientImpl struct {
	cc *http.Client
}

func NewTreeHTTPClient(client *http.Client) TreeHTTPClient {
	return &TreeHTTPClientImpl{client}
}

func (c *TreeHTTPClientImpl) GetTree(ctx context.Context, in *GetTreeRequest, opts ...http.CallOption) (*GetTreeReply, error) {
	var out GetTreeReply
	pattern := "/gettree/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTreeGetTree))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
