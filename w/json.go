package w

import (
	"encoding/json"
)

var Json *w_json

func init() {
	Json = &w_json{}
}

// 打印漂亮的json
func (j *w_json) Json_print(data interface{}) (jsonString string) {
	jsonByte, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(jsonByte)
}

// json序列化
func (j *w_json) Json_dumps(data interface{}) (jsonString string, err error) {
	jsonByte, err := json.Marshal(data)
	if err != nil {
		return jsonString, err
	}
	return string(jsonByte), err
}

// json序列化
func (j *w_json) Json_dump_ident(data interface{}) (jsonString string, err error) {
	jsonByte, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	jsonString = string(jsonByte)
	return
}
