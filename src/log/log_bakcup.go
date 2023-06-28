package log

import (
	"errors"
	"fmt"
	"github.com/wanggaolin/go_lib/src/file"
	"github.com/wanggaolin/go_lib/src/lib"
	"io"
	"os"
	"path"
	"regexp"
	"strconv"
	"time"
)

// date: 2021/12/10
// email: brach@lssin.com

func Interface_log_backup(file_path string, back_dir string, file_action string) (bool, error) {
	if file.PathExists(file_path) == false {
		return false, errors.New(file_path + " :No such file or directory")
	}

	_t := time.Now()
	backup_directory := path.Join(back_dir, _t.Format("2006"), _t.Format("2006-01"), _t.Format("2006-01-02"))

	reg_filter := [...]string{
		"_(20\\d{2})-(\\d{1,2})-(\\d{1,2})_",
		"-(20\\d{2})-(\\d{1,2})-(\\d{1,2})-",
		"-(20\\d{2})-(\\d{1,2})-(\\d{1,2})_",
		"_(20\\d{2})-(\\d{1,2})-(\\d{1,2})-",
	}

	for _, reg := range reg_filter {
		filter_time := regexp.MustCompile(reg).FindStringSubmatch(file_path)
		if len(filter_time) == 4 {
			month, _err_month := strconv.Atoi(filter_time[2])
			day, _err_day := strconv.Atoi(filter_time[3])
			if _err_month != nil {
				continue
			}
			if _err_day != nil {
				continue
			}
			if day > 30 || month > 12 {
				continue
			}
			backup_directory = path.Join(back_dir, filter_time[1], filter_time[1]+"-"+filter_time[2], filter_time[1]+"-"+filter_time[2]+"-"+filter_time[3])
			break
		}
	}

	if file.PathExists(backup_directory) == false {
		_error := os.MkdirAll(backup_directory, os.ModePerm)
		if _error != nil {
			return false, _error
		}
	}

	if file.IsDir(backup_directory) == false {
		return false, errors.New(backup_directory + " :This directory is not a directory")
	}
	new_file := path.Join(backup_directory, path.Base(file_path))

	if file.PathExists(new_file) {
		Log_warr(new_file, ":This file already exists")
		for _, n := range lib.Make_range(1, 1024) {
			if file.PathExists(fmt.Sprintf("%v.%v", new_file, n)) == false {
				new_file = fmt.Sprintf("%v.%v", new_file, n)
				break
			}
			if n > 500 {
				return false, errors.New("Too many duplicate file names")
			}
		}
	}

	destFile, err := os.OpenFile(new_file, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return false, err
	}
	srcFile, err := os.Open(file_path)
	if err != nil {
		return false, err
	}
	nBytes, err := io.Copy(destFile, srcFile)
	if err != nil {
		return false, err
	}
	_ = nBytes
	if file_action == "mv" {
		_err := os.Remove(file_path)
		if _err != nil {
			return false, _err
		}
		Log_debug(file_path, new_file, "Move success")
	} else {
		Log_debug(file_path, new_file, "Copy success")
	}
	return true, nil
}
