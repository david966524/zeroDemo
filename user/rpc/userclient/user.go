// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"

	"zero/user/rpc/types/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	IdRequest    = user.IdRequest
	ResponseMsg  = user.ResponseMsg
	UserInfo     = user.UserInfo
	UserResponse = user.UserResponse

	User interface {
		GetUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserResponse, error)
		AddUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*ResponseMsg, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) GetUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}

func (m *defaultUser) AddUser(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*ResponseMsg, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.AddUser(ctx, in, opts...)
}