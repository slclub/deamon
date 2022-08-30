package client

import (
	"github.com/slclub/deamon/plugin/rpcx"
)

type Client struct {
}

func New() *Client {
	rt := &Client{}
	//rt.RPCX().Register("DockerServer")
	return rt
}

func (this *Client) RPCX() *rpcx.RpcxClient {
	return rpcx.Default
}

func (this *Client) Init(fnn func()) {
	fnn()
	rpcx.RequestRegister("DockerServer")
}
