package rpcx

import (
	"context"
	"github.com/smallnest/rpcx/client"
	"math/rand"
	"strings"
	"time"
)

const SERVER_ID_START = 10000

/**
 *
 */
type RpcxClient struct {
	Len        int
	mapclients map[string][]*agentClient
	tokenAddr  map[string]string // 玩家token对应的 游戏服务地址
}

var ServiceAddrs []string = []string{}

var Default = NewRpcxClient()
var RequestRegister func(controller string)

func init() {
	RequestRegister = Default.Register
	//RequestRegister("Arith") // example
}

func NewRpcxClient() *RpcxClient {
	return &RpcxClient{
		Len: 0,
		//xclientList: make([]*agentClient, 0, 10),
		tokenAddr:  make(map[string]string),
		mapclients: make(map[string][]*agentClient),
	}
}
func (rp *RpcxClient) Register(controller string) {
	//clients_conf = utils.Config("client").GetStringSlice("SrvsDeamon")
	if len(ServiceAddrs) == 0 {
		return
	}

	//discover, _ := client.NewMultipleServersDiscovery(addrPairs)
	for i, addr := range ServiceAddrs {
		if strings.Trim(addr, " ") == "" {
			continue
		}
		if len(rp.mapclients[controller]) == 0 {
			rp.mapclients[controller] = []*agentClient{}
		}
		rp.mapclients[controller] = append(rp.mapclients[controller], &agentClient{})
		rp.mapclients[controller][i].InitStart(controller, addr)
		rp.mapclients[controller][i].initIndex(i)
		rp.Len++
	}
	//fmt.Println("[rpcx][getClients]", controller, "all:", rp.xclientList)
}

func (rp *RpcxClient) GetRand(controller string) *agentClient {
	mclients := rp.GetClientByController(controller)
	rand_index, ll := 0, len(mclients)
	if ll > 1 {
		rand_index = rand.Intn(ll)
	}

	return mclients[rand_index]
}

func (rp *RpcxClient) GetByAddr(controller string, addr string) *agentClient {
	clients := rp.mapclients[controller]
	for _, cli := range clients {
		if cli.GetAddr() == addr {
			return cli
		}
	}
	return nil
}

// 此方法暂时不用，token 控制不好，有泄漏内存的风险
func (rp *RpcxClient) GetClientByToken(controller string, token string) *agentClient {
	addr, ok := rp.tokenAddr[token]
	if !ok {
		return nil
	}
	return rp.GetByAddr(controller, addr)
}

func (rp *RpcxClient) CallSlice(controller string, method string, fn func(aclient *agentClient, method string)) {
	mclients := rp.GetClientByController(controller)
	for i, n := 0, len(mclients); i < n; i++ {
		fn(mclients[i], method)
	}
}

func (rp *RpcxClient) GetClientByController(controller string) []*agentClient {
	return rp.mapclients[controller]
}

func (rp *RpcxClient) Stop() {
	for _, clis := range rp.mapclients {
		for _, cli := range clis {
			cli.Close()
		}
	}
}

// 根据服务ID ，索引ID  获取对应的agent client
// 用于获取服务信息
func (rp *RpcxClient) GetAgentClientByIndex(id int) []*agentClient {
	rtn := []*agentClient{}
	for _, clis := range rp.mapclients {
		for _, cli := range clis {
			if cli.GetIndex() == id {
				rtn = append(rtn, cli)
				break
			}
		}
	}
	return nil
}

// ===========================================================================
type agentClient struct {
	controller string
	ip         string
	port       string
	index_id   int // 可以作为服务ID
	xclient    client.XClient
	connected  bool
}

func (ac *agentClient) InitStart(controller string, addr string) {
	addr_arr := strings.Split(addr, ":")
	ac.ip = addr_arr[0]
	ac.port = addr_arr[1]

	ac.controller = controller
	discover, _ := client.NewPeer2PeerDiscovery("tcp@"+addr, "")
	option := client.DefaultOption
	option.Heartbeat = true
	option.HeartbeatInterval = time.Second * 5
	ac.xclient = client.NewXClient(controller, client.Failtry, client.RandomSelect, discover, option)
}

func (ac *agentClient) Call(ctx context.Context, method string, args interface{}, reply interface{}) error {
	err := ac.xclient.Call(ctx, method, args, reply)
	if err != nil {
		ac.connected = false
	} else {
		if !ac.Connected() {
			//ac.ConfirmeServerSync(ctx)
		}
		ac.connected = true
	}
	return err
}

func (ac *agentClient) ConfirmeServerSync(ctx context.Context) {
	acs := Default.GetClientByController("ServerDescSync")
	if len(acs) < ac.GetIndex() {
		return
	}
	var tac *agentClient = acs[ac.GetIndex()]

	// 判空
	if tac.GetPort() == "" {
		return
	}

	server_id := tac.GetIndex() + SERVER_ID_START
	req := &struct {
		ID int
	}{ID: server_id}
	reply := &struct {
		MsgCode int
	}{}
	tac.xclient.Call(ctx, "SyncID", req, reply)
	return
}

func (ac *agentClient) Close() {
	ac.xclient.Close()
}

func (ac *agentClient) GetAddr() string {
	return strings.Join([]string{ac.ip, ac.port}, ":")
}

func (ac *agentClient) GetIp() string {
	return ac.ip
}

func (ac *agentClient) GetPort() string {
	return ac.port
}

func (ac *agentClient) initIndex(i int) {
	ac.index_id = i
}

func (ac *agentClient) GetIndex() int {
	return ac.index_id
}

func (ac *agentClient) GetGameListenAddr() string {
	//ServiceAddrs := utils.Config("client").GetStringSlice("SrvsDeamon")
	if len(ServiceAddrs)-1 < ac.index_id {
		return ""
	}
	return ServiceAddrs[ac.index_id]
}

func (ac *agentClient) Connected() bool {
	return ac.connected
}
