package file

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
)

// date: 2022/02/17
// email: brach@lssin.com

func File_copy(src_path string, dst_path string) (bool, error) {
	srcFile, err := os.Open(src_path)
	if err != nil {
		return false, err
	}
	destFile, err := os.OpenFile(dst_path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return false, err
	}
	nBytes, err := io.Copy(destFile, srcFile)
	_ = nBytes
	if err != nil {
		return false, err
	}
	return true, nil
}

func ReadFile(file_name string) (string, error) {
	jsonFile, err := os.Open(file_name)
	if err != nil {
		return "", err
	}
	defer jsonFile.Close()
	byteValue, _err := ioutil.ReadAll(jsonFile)
	if _err != nil {
		return "", _err
	}
	return string(byteValue), nil
}

func FileContent_to_Map(file_path string) (map[string]interface{}, error) {
	file_content, file_error := ReadFile(file_path)
	if file_error != nil {
		return nil, file_error
	}
	var resMap map[string]interface{}
	json_err := json.Unmarshal([]byte(file_content), &resMap)
	if json_err != nil {
		return nil, json_err
	}
	return resMap, nil
}

func JsonFile_to_yaml(file_path string) (map[string]interface{}, error) {
	t := map[string]interface{}{}
	buffer, err := ioutil.ReadFile(file_path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(buffer, &t)
	if err != nil {
		panic(err.Error())
	}
	return t, nil
}
