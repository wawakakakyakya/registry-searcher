package requester

import (
	"registry-searcher/registry_url"
)

type GetTagRequester struct {
	// url URL
	url registry_url.RegistryUrlInterface
	*baseRequester
}

func (r *GetTagRequester) GetUrl() (string, error) {
	url, err := r.url.MakeUrl()
	if err != nil {
		return "", err
	}
	return url, err
}

func (r *GetTagRequester) Execute(url string) (string, error) {
	return r.get(url)
}

func NewGetTagRequester(url registry_url.RegistryUrlInterface) RequesterInterface {
	return &GetTagRequester{url: url}
}
