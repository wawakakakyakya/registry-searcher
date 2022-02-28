package registry_url

import (
	"fmt"
	"net/url"
	"strings"
)

type GetManifestUrl struct {
	Address string
	Port    int
	Version int
	UseSsl  bool
	Image   string
	*baseRegistryURL
}

func (l *GetManifestUrl) makeAbsURL() string {
	baseUrl := fmt.Sprintf("%s://%s", l.getScheme(l.UseSsl), l.getAuthority(l.Address, l.Port))
	imageTag := strings.Split(l.Image, ":")
	if len(imageTag) == 1 {
		imageTag = append(imageTag, "latest")
	}
	return fmt.Sprintf("%s/%s/%s/manifests/%s", baseUrl, l.getVersion(l.Version), imageTag[0], imageTag[1])
}

//entry point
func (l *GetManifestUrl) MakeUrl() (string, error) {
	absUrl := l.makeAbsURL()
	if _, err := url.Parse(absUrl); err != nil {
		return "", err
	}
	return absUrl, nil
}
