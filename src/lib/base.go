package lib

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// date: 2021/12/10
// email: brach@lssin.com

func Make_range(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func Show_version(GoVersion string, Auchar string, AppVersion string) {
	fmt.Fprintf(os.Stderr, `App Name: %v
App Auchar: %v
GoLang Version: %v
`, AppVersion, Auchar, GoVersion)
}

func Check_ip(ip_name string) bool {
	re_ip := regexp.MustCompile(`^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})$`)
	filter_ip := re_ip.FindStringSubmatch(ip_name)
	if len(filter_ip) > 1 {
		a, _ := strconv.Atoi(filter_ip[1])
		b, _ := strconv.Atoi(filter_ip[2])
		c, _ := strconv.Atoi(filter_ip[3])
		d, _ := strconv.Atoi(filter_ip[4])
		if a <= 255 && b <= 255 && c <= 255 && d <= 255 && a > 0 && d > 0 {
			return true
		}
	}
	return false
}

func List_file(directory_name string) ([]string, error) {
	var s []string
	return _list_dir(directory_name, s)
}

func List_directory(directory_name string) ([]fs.FileInfo, error) {
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

func FileState(file_path string) (os.FileInfo, error) {
	f, err := os.Open(file_path)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	fi, errs := f.Stat()
	if errs != nil {
		return nil, errs
	}
	return fi, nil
}

func Unix_for_kb1(unit_size int64) string {
	if unit_size > 1024 {
		k := float64(unit_size) / 1024
		if k > 1024 {
			m := k / 1024
			if m > 1024 {
				return fmt.Sprintf("%.1fG", m/1024)
			}
			return fmt.Sprintf("%.1fM", m)
		}
		return fmt.Sprintf("%.1fK", k)
	}
	return fmt.Sprintf("%v", unit_size)
}

func Shell_run(command string, arg ...string) (string, error) {
	cmd := exec.Command(command, arg...)
	cmd.Env = os.Environ()
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", errors.New(fmt.Sprintf("%v(%v)", strings.TrimSpace(stderr.String()), err.Error()))
	}
	return out.String(), nil
}
