package requester

import (
	"fmt"
	"registry-searcher/registry_url"
)

type DeleteImageRequester struct {
	// url URL
	url *registry_url.DeleteImageUrl
	*baseRequester
}

func (r *DeleteImageRequester) GetUrl() (string, error) {
	url, err := r.url.MakeUrl()
	if err != nil {
		return "", err
	}
	return url, err
}

func (r *DeleteImageRequester) delete() (string, error) {
	manifestUrl := registry_url.GetManifestUrl{}
	manifestUrl.Address = r.url.Address
	manifestUrl.Port = r.url.Port
	manifestUrl.Version = r.url.Version
	manifestUrl.UseSsl = r.url.UseSsl
	manifestUrl.Image = r.url.Image
	getManifestRequester := NewGetManifestRequester(&manifestUrl)
	manifestAbsUrl, err := getManifestRequester.GetUrl()
	fmt.Println(manifestAbsUrl)
	if err != nil {
		return "", err
	}
	h := header{key: "Accept", value: "application/vnd.docker.distribution.manifest.v2+json"}

	getResp, err := r.do("GET", manifestAbsUrl, headers{h})
	if err != nil {
		fmt.Println("get digest failed")
		return "", err
	}

	defer getResp.Body.Close()
	r.url.Digest = getResp.Header["Docker-Content-Digest"][0]

	absUrl, err := r.url.MakeUrl()
	if err != nil {
		return "", err
	}

	delResp, err := r.do("DELETE", absUrl, headers{})
	if err != nil {
		fmt.Println("delete image failed")
		return "", err
	}
	defer delResp.Body.Close()

	return printBody(delResp.Body), nil
}

func (r *DeleteImageRequester) Execute(url string) (string, error) {
	return r.delete()
}

func NewDeleteImageRequester(url *registry_url.DeleteImageUrl) RequesterInterface {
	return &DeleteImageRequester{url: url}
}
