package msg

type Request struct {
	ServID int
	Name   string
	Data   map[string]any
}

type Response struct {
	MsgCode int
	Data    any
	Desc    string
}

type ServItem struct {
	Index int
	Name  string
	State string // 链接状态
	Addr  string
}
