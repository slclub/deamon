package model

import "github.com/slclub/deamon/helper"

type DeamonProcceser struct {
	ID    int
	Name  string
	State string // enable / disable
}

type deamonBox struct {
	data []*DeamonProcceser
}

var _deamon deamonBox = deamonBox{data: []*DeamonProcceser{}}

func Mgr() *deamonBox {
	return &_deamon
}

func (this *deamonBox) Add(dea *DeamonProcceser) {
	for i, da := range this.data {
		if da.Name == dea.Name {
			this.data[i] = dea
			return
		}
	}
	this.data = append(this.data, dea)
}

func (this *deamonBox) GetByName(name string) *DeamonProcceser {
	for _, da := range this.data {
		if da.Name == name {
			return da
		}
	}
	return nil
}

func (this *deamonBox) GetByID(id int) *DeamonProcceser {
	for _, da := range this.data {
		if da.ID == id {
			return da
		}
	}
	return nil
}

func (this *deamonBox) Check(name string) bool {
	dea := this.GetByName(name)
	if dea == nil {
		return true
	}
	if dea.State == helper.DEAMON_ENABLE {
		return true
	}
	return false
}
