package requester

import (
	"io"
	"io/ioutil"
	"registry-searcher/registry_url"
)

func printBody(body io.ReadCloser) string {
	_body, _ := ioutil.ReadAll(body)
	// body は []byte
	//fmt.Println(_body)
	return string(_body)
}

type ListImageRequester struct {
	// url URL
	url registry_url.RegistryUrlInterface
	*baseRequester
}

func (l *ListImageRequester) GetUrl() (string, error) {
	url, err := l.url.MakeUrl()
	if err != nil {
		return "", err
	}
	return url, err
}

func NewListImageRequester(url registry_url.RegistryUrlInterface) RequesterInterface {
	return &ListImageRequester{url: url}
}
