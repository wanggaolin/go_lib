package pub

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// date: 2022/01/05
// email: brach@lssin.com

var Request *requests

func init() {
	Request = &requests{}
}

type requests struct {
}

func (r requests) Post(host string, post string, timeout_size int64) (string, int, error) {
	timeout := time.Duration(timeout_size) * time.Second
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Post(host, "application/x-www-form-urlencoded", strings.NewReader(post))
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, err
	}
	return string(body), resp.StatusCode, nil
}
