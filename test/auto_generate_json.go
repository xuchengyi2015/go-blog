package main

import (
	"encoding/json"
	"strings"
)

func main() {

}

func MyMarshal(obj interface{}) (b []byte, e error) {
	b, e = json.Marshal(obj)
	if e != nil {
		return
	}
	var m map[string]interface{}
	e = json.Unmarshal(b, &m)
	if e != nil {
		return
	}
	HandleMapStyle(m)
	return json.Marshal(m)
}

func HandleMapStyle(m map[string]interface{}) {
	for key, value := range m {
		switch v := value.(type) {
		case []interface{}:
			for _, vv := range v {
				if elem, ok := vv.(map[string]interface{}); ok {
					HandleMapStyle(elem)
				}
			}
		case map[string]interface{}:
			HandleMapStyle(v)
		}
		delete(m, key)
		m[strings.ToLower(key)] = value //此处简化处理, 全变小写
	}
}
