package registry_url

import (
	"fmt"
	"net/url"
)

type GetManifestUrl struct {
	Address string
	Port    int
	Version int
	UseSsl  bool
	Tag     string
	Image   string
	*baseRegistryURL
}

func (l *GetManifestUrl) makeAbsURL() string {
	baseUrl := fmt.Sprintf("%s://%s", l.getScheme(l.UseSsl), l.getAuthority(l.Address, l.Port))
	return fmt.Sprintf("%s/%s/%s/manifests/%s", baseUrl, l.getVersion(l.Version), l.Image, l.Tag)
}

//entry point
func (l *GetManifestUrl) MakeUrl() (string, error) {
	absUrl := l.makeAbsURL()
	if _, err := url.Parse(absUrl); err != nil {
		return "", err
	}
	return absUrl, nil
}
