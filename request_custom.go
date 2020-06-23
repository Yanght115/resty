package resty

import (
	"encoding/json"
	"fmt"
)

func checkResponseStatus(resp *Response) (err error) {
	status := resp.StatusCode()
	if status < 200 || status > 299 {
		err = fmt.Errorf("get response with status:%d", status)
	}
	return
}

// GetObject will unmarshal get response content to obj, response status must be 20x
func (r *Request) GetToObject(url string, obj interface{}) (resp *Response, err error) {
	if resp, err = r.Execute(MethodGet, url); err != nil {
		return
	}
	if err = checkResponseStatus(resp); err == nil {
		err = json.Unmarshal(resp.Body(), obj)
	}
	return
}

// PostObject will unmarshal post response content to obj, response status must be 20x
func (r *Request) PostToObject(url string, obj interface{}) (resp *Response, err error) {
	if resp, err = r.Execute(MethodPost, url); err != nil {
		return
	}
	if err = checkResponseStatus(resp); err == nil {
		err = json.Unmarshal(resp.Body(), obj)
	}
	return
}
