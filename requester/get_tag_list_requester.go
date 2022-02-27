package requester

import (
	"registry-searcher/registry_url"
)

type GetTagRequester struct {
	// url URL
	url registry_url.RegistryUrlInterface
	*baseRequester
}

func (l *GetTagRequester) GetUrl() (string, error) {
	url, err := l.url.MakeUrl()
	if err != nil {
		return "", err
	}
	return url, err
}

func NewGetTagRequester(url registry_url.RegistryUrlInterface) RequesterInterface {
	return &GetTagRequester{url: url}
}
