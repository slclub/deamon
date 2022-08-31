package helper

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func ConvAnyToInt(value any) int {
	switch val := value.(type) {
	case string:
		rtn, err := strconv.Atoi(val)
		if err != nil {
			return 0
		}
		return rtn
	case int:
		return val
	case int64:
		return int(val)
	case int32:
		return int(val)
	case uint64:
		return int(val)
	case uint32:
		return int(val)
	case float64:
		return int(val)
	}
	return 0
}

func ConvSizeByteToFormat(size_in any) string {
	var size_m float32 = 1000000
	var size_k float32 = 1000
	var size float32 = 0
	switch val := size_in.(type) {
	case int:
		size = float32(val)
	case int32:
		size = float32(val)
	case int64:
		size = float32(val)
	case uint:
		size = float32(val)
	}
	sep_unit := ""
	switch {
	case size > size_m:
		size = size / size_m
		sep_unit = "M"
	case size > size_k && size <= size_m:
		size = size / size_k
		sep_unit = "K"
	default:
		sep_unit = "B"
	}
	return fmt.Sprintf("%.2f%s", size, sep_unit)
}

func ConvArrFloat64ToInt(fa []any) []int {
	rtn := make([]int, 0, len(fa))
	for _, v := range fa {
		rtn = append(rtn, ConvAnyToInt(v))
	}
	return rtn
}

func ConvStruct2Map(v any) map[string]any {
	data, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	json_data := make(map[string]any)
	json.Unmarshal(data, &json_data)
	return json_data
}
