package requester

import (
	"errors"
	"fmt"
	"net/http"
)

type header struct {
	key   string
	value string
}
type headers []header

type baseRequester struct {
}

func (r *baseRequester) do(method string, url string, headers headers) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	for _, h := range headers {
		req.Header.Set(h.key, h.value)
	}
	if err != nil {
		return nil, err
	}
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	ok := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !ok {
		fmt.Printf("status code(%d) was not succeeded\n", resp.StatusCode)
		return nil, errors.New(printBody(resp.Body))
	}
	return resp, nil
}

func (r *baseRequester) get(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ok := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !ok {
		fmt.Printf("status code(%d) was not succeeded\n", resp.StatusCode)
		return "", errors.New(printBody(resp.Body))
	}

	return printBody(resp.Body), nil
}
