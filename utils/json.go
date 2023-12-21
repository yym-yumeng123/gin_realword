package utils

import "encoding/json"

// JsonMarshal 返回 json 字符串
func JsonMarshal(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}
