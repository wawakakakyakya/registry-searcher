package requester

import (
	"registry-searcher/registry_url"
)

type GetManifestRequester struct {
	// url URL
	url registry_url.RegistryUrlInterface
	*baseRequester
}

func (r *GetManifestRequester) GetUrl() (string, error) {
	url, err := r.url.MakeUrl()
	if err != nil {
		return "", err
	}
	return url, err
}

func (r *GetManifestRequester) Execute(url string) (string, error) {
	return r.get(url)
}

func NewGetManifestRequester(url registry_url.RegistryUrlInterface) RequesterInterface {
	return &GetManifestRequester{url: url}
}
