package main

import (
	"fmt"
	"github.com/slclub/deamon/client"
	"github.com/slclub/deamon/client/servsys"
	"github.com/slclub/deamon/plugin/rpcx"
	"github.com/slclub/deamon/utils"
)

func main() {
	cclient := client.New()
	cclient.Init(func() {
		rpcx.ServiceAddrs = utils.Config("client").GetStringSlice("SrvsDeamon")
	})
	//vdatas := servsys.GetServListByName(cclient.RPCX(), "tiandi-meta", "DockerServer")
	vdatas := servsys.GetServListByName(cclient.RPCX(), "tiandi-meta")
	fmt.Println("service list data :", *vdatas[0])

	servsys.StartWithName(cclient.RPCX(), "tiandi-meta", rpcx.ServiceAddrs[0])
}
