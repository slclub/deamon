package utils

import "reflect"

// 判断空接口
func IsNil(i interface{}) bool {
	defer func() {
		recover()
	}()
	vi := reflect.ValueOf(i)
	return vi.IsNil()
}
