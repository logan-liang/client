// Code generated by goctl. DO NOT EDIT!
// Source: test.proto

package client

import (
	"context"

	"github.com/logan-liang/test/test"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	Request  = test.Request
	Response = test.Response

	Test interface {
		Ping(ctx context.Context, in *Request) (*Response, error)
	}

	defaultTest struct {
		cli zrpc.Client
	}
)

func NewTest(cli zrpc.Client) Test {
	return &defaultTest{
		cli: cli,
	}
}

func (m *defaultTest) Ping(ctx context.Context, in *Request) (*Response, error) {
	client := test.NewTestClient(m.cli.Conn())
	return client.Ping(ctx, in)
}
