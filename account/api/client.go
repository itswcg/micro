package api

import (
	"context"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"google.golang.org/grpc"
)

const AppId = "account.service"
const target = "direct://default/127.0.0.1:9001"

func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (AccountClient, error) {
	client := warden.NewClient(cfg, opts...)
	conn, err := client.Dial(context.Background(), target)
	if err != nil {
		return nil, err
	}
	return NewAccountClient(conn), nil
}
