package requester

import (
	"registry-searcher/registry_url"
)

type GetManifestRequester struct {
	// url URL
	url registry_url.RegistryUrlInterface
	*baseRequester
}

func (l *GetManifestRequester) GetUrl() (string, error) {
	url, err := l.url.MakeUrl()
	if err != nil {
		return "", err
	}
	return url, err
}

func NewGetManifestRequester(url registry_url.RegistryUrlInterface) RequesterInterface {
	return &GetManifestRequester{url: url}
}
