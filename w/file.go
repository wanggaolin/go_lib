package w

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
)

var File *fileing

// date: 2022/02/17
// email: brach@lssin.com

func init() {
	File = &fileing{}
}
func (f *fileing) File_copy(src_path string, dst_path string) (bool, error) {
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

func (f *fileing) ReadFile(file_name string) (string, error) {
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

func (f *fileing) FileContent_to_Map(file_path string) (map[string]interface{}, error) {
	file_content, file_error := f.ReadFile(file_path)
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

func (f *fileing) JsonFile_to_yaml(file_path string) (map[string]interface{}, error) {
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

func (f *fileing) FileState(file_path string) (os.FileInfo, error) {
	file, err := os.Open(file_path)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	fi, errs := file.Stat()
	if errs != nil {
		return nil, errs
	}
	return fi, nil
}
func (f *fileing) PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func (f *fileing) IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func _list_dir(filePath string, s []string) ([]string, error) {
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		return nil, err
	} else {
		for _, v := range files {
			if v.IsDir() {
				s, _ = _list_dir(filePath+"/"+v.Name(), s)
			} else {
				s = append(s, filePath+"/"+v.Name())
			}
		}
	}
	return s, nil
}

func (f *fileing) List_file(directory_name string) ([]string, error) {
	var s []string
	return _list_dir(directory_name, s)
}

func (f *fileing) List_directory(directory_name string) ([]fs.FileInfo, error) {
	var data []fs.FileInfo
	file_list, _error := ioutil.ReadDir(directory_name)
	if _error != nil {
		return nil, _error
	}
	for _, file_name := range file_list {
		data = append(data, file_name)
	}
	if data == nil {
		data = append(data)
	}
	return data, nil
}

func (f *fileing) File_md5(file_path string) (md5String string, err error) {
	file, err := os.Open(file_path)
	if err != nil {
		return md5String, err
	}
	defer file.Close()

	file_hash := md5.New()
	if _, err = io.Copy(file_hash, file); err != nil {
		return md5String, err
	}
	md5Hash := file_hash.Sum(nil)
	md5String = hex.EncodeToString(md5Hash)
	return md5String, err
}
