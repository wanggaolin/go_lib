package w

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

func (r *requests) Post(host string, post string, timeout_size int64) (string, int, error) {
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

func (r *requests) Head(host string, timeout_size int64) (header *http.Response, err error) {
	timeout := time.Duration(timeout_size) * time.Second
	client := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("HEAD", host, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36")
	if err != nil {
		return header, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return header, err
	}
	defer resp.Body.Close()
	return resp, err
}
