package requester

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

// handler ...
type handler struct{}

// New ...
func New() Contract {
	return &handler{}
}

// RAW ...
func (request *handler) RAW(method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, url, body)
}

// GET request type get
func (request *handler) GET(url string, header map[string]string) ([]byte, error) {
	var result []byte
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}
	if header != nil {
		for content, value := range header {
			req.Header.Set(content, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	return body, nil
}

// POST request type post
func (request *handler) POST(url string, header map[string]string, payload []byte) ([]byte, error) {
	var result []byte
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return result, err
	}
	if header != nil {
		for content, value := range header {
			req.Header.Set(content, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	return body, nil
}

// PUT request type post
func (request *handler) PUT(url string, header map[string]string, payload []byte) ([]byte, error) {
	var result []byte
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payload))
	if err != nil {
		return result, err
	}
	if header != nil {
		for content, value := range header {
			req.Header.Set(content, value)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	return body, nil
}

// DELETE request type get
func (request *handler) DELETE(url string, header map[string]string) ([]byte, error) {
	var result []byte
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return result, err
	}
	// set default headers
	req.Header.Add("Content-Type", "application/json")
	// add optional headers
	for content, value := range header {
		req.Header.Set(content, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	return body, nil
}
