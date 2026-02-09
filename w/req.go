package w

import (
	"bytes"
	"io"
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

// post from
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

// post Json
type Requests_body struct {
	Timeout int64
	Body    []byte // post json
	Header  map[string]string
}

func (r *requests) Post_v1(host string, http_body *Requests_body) (string, int, error) {
	timeout := time.Duration(http_body.Timeout) * time.Second
	client := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("POST", host, bytes.NewBuffer(http_body.Body))
	if err != nil {
		return "", 0, err
	}
	for k, v := range http_body.Header {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", 0, err
	}
	return string(body), resp.StatusCode, nil
}

func (r *requests) Head(host string, http_body *Requests_body) (header *http.Response, err error) {
	timeout := time.Duration(http_body.Timeout) * time.Second
	client := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("HEAD", host, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36")
	for k, v := range http_body.Header {
		req.Header.Set(k, v)
	}
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

func (r *requests) Get(host string, http_body *Requests_body) (response string, code int, err error) {
	timeout := time.Duration(http_body.Timeout) * time.Second
	client := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("GET", host, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36")
	for k, v := range http_body.Header {
		req.Header.Set(k, v)
	}
	if err != nil {
		return response, code, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return response, code, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, code, err
	}
	return string(body), resp.StatusCode, nil
}
