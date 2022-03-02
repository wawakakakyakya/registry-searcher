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

func (r *ListImageRequester) GetUrl() (string, error) {
	url, err := r.url.MakeUrl()
	if err != nil {
		return "", err
	}
	return url, err
}

func (r *ListImageRequester) Execute(url string) (string, error) {
	return r.get(url)
}

func NewListImageRequester(url registry_url.RegistryUrlInterface) RequesterInterface {
	return &ListImageRequester{url: url}
}
