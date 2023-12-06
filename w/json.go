package w

import (
	"encoding/json"
)

var Json *w_json

func init() {
	Json = &w_json{}
}

func (j *w_json) Json_print(data interface{}) (jsonString string, err error) {
	jsonByte, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return jsonString, err
	}
	return string(jsonByte), err
}

func (j *w_json) Json_dumps(data interface{}) (jsonString string, err error) {
	jsonByte, err := json.Marshal(data)
	if err != nil {
		return jsonString, err
	}
	return string(jsonByte), err
}
