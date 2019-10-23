package api

import (
	"context"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"google.golang.org/grpc"
)

const AppId = "leaf.service"

func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (LeafClient, error) {
	client := warden.NewClient(cfg, opts...)
	conn, err := client.Dial(context.Background(), "discovery://default/"+AppId)
	if err != nil {
		return nil, err
	}
	return NewLeafClient(conn), nil
}
