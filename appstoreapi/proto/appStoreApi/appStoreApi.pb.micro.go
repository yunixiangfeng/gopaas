// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/appStoreApi/appStoreApi.proto

package appStoreApi

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for AppStoreApi service

func NewAppStoreApiEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AppStoreApi service

type AppStoreApiService interface {
	FindAppStoreById(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	AddAppStore(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	DeleteAppStoreById(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	UpdateAppStore(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	//默认接口
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	//添加常用的统计接口
	AddInstallNum(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetInstallNum(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	AddViewNum(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetViewNum(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type appStoreApiService struct {
	c    client.Client
	name string
}

func NewAppStoreApiService(name string, c client.Client) AppStoreApiService {
	return &appStoreApiService{
		c:    c,
		name: name,
	}
}

func (c *appStoreApiService) FindAppStoreById(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AppStoreApi.FindAppStoreById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appStoreApiService) AddAppStore(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AppStoreApi.AddAppStore", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appStoreApiService) DeleteAppStoreById(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AppStoreApi.DeleteAppStoreById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appStoreApiService) UpdateAppStore(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AppStoreApi.UpdateAppStore", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appStoreApiService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AppStoreApi.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appStoreApiService) AddInstallNum(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AppStoreApi.AddInstallNum", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appStoreApiService) GetInstallNum(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AppStoreApi.GetInstallNum", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appStoreApiService) AddViewNum(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AppStoreApi.AddViewNum", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appStoreApiService) GetViewNum(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AppStoreApi.GetViewNum", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AppStoreApi service

type AppStoreApiHandler interface {
	FindAppStoreById(context.Context, *Request, *Response) error
	AddAppStore(context.Context, *Request, *Response) error
	DeleteAppStoreById(context.Context, *Request, *Response) error
	UpdateAppStore(context.Context, *Request, *Response) error
	//默认接口
	Call(context.Context, *Request, *Response) error
	//添加常用的统计接口
	AddInstallNum(context.Context, *Request, *Response) error
	GetInstallNum(context.Context, *Request, *Response) error
	AddViewNum(context.Context, *Request, *Response) error
	GetViewNum(context.Context, *Request, *Response) error
}

func RegisterAppStoreApiHandler(s server.Server, hdlr AppStoreApiHandler, opts ...server.HandlerOption) error {
	type appStoreApi interface {
		FindAppStoreById(ctx context.Context, in *Request, out *Response) error
		AddAppStore(ctx context.Context, in *Request, out *Response) error
		DeleteAppStoreById(ctx context.Context, in *Request, out *Response) error
		UpdateAppStore(ctx context.Context, in *Request, out *Response) error
		Call(ctx context.Context, in *Request, out *Response) error
		AddInstallNum(ctx context.Context, in *Request, out *Response) error
		GetInstallNum(ctx context.Context, in *Request, out *Response) error
		AddViewNum(ctx context.Context, in *Request, out *Response) error
		GetViewNum(ctx context.Context, in *Request, out *Response) error
	}
	type AppStoreApi struct {
		appStoreApi
	}
	h := &appStoreApiHandler{hdlr}
	return s.Handle(s.NewHandler(&AppStoreApi{h}, opts...))
}

type appStoreApiHandler struct {
	AppStoreApiHandler
}

func (h *appStoreApiHandler) FindAppStoreById(ctx context.Context, in *Request, out *Response) error {
	return h.AppStoreApiHandler.FindAppStoreById(ctx, in, out)
}

func (h *appStoreApiHandler) AddAppStore(ctx context.Context, in *Request, out *Response) error {
	return h.AppStoreApiHandler.AddAppStore(ctx, in, out)
}

func (h *appStoreApiHandler) DeleteAppStoreById(ctx context.Context, in *Request, out *Response) error {
	return h.AppStoreApiHandler.DeleteAppStoreById(ctx, in, out)
}

func (h *appStoreApiHandler) UpdateAppStore(ctx context.Context, in *Request, out *Response) error {
	return h.AppStoreApiHandler.UpdateAppStore(ctx, in, out)
}

func (h *appStoreApiHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.AppStoreApiHandler.Call(ctx, in, out)
}

func (h *appStoreApiHandler) AddInstallNum(ctx context.Context, in *Request, out *Response) error {
	return h.AppStoreApiHandler.AddInstallNum(ctx, in, out)
}

func (h *appStoreApiHandler) GetInstallNum(ctx context.Context, in *Request, out *Response) error {
	return h.AppStoreApiHandler.GetInstallNum(ctx, in, out)
}

func (h *appStoreApiHandler) AddViewNum(ctx context.Context, in *Request, out *Response) error {
	return h.AppStoreApiHandler.AddViewNum(ctx, in, out)
}

func (h *appStoreApiHandler) GetViewNum(ctx context.Context, in *Request, out *Response) error {
	return h.AppStoreApiHandler.GetViewNum(ctx, in, out)
}
