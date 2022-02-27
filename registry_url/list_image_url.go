package registry_url

import (
	"fmt"
	"net/url"
)

type ListImageUrl struct {
	Address string
	Port    int
	Version int
	UseSsl  bool
	*baseRegistryURL
}

//コマンドごとに異なる
func (l *ListImageUrl) makeAbsURL() string {
	baseUrl := fmt.Sprintf("%s://%s", l.getScheme(l.UseSsl), l.getAuthority(l.Address, l.Port))
	return fmt.Sprintf("%s/%s/_catalog", baseUrl, l.getVersion(l.Version))
}

//entry point
func (l *ListImageUrl) MakeUrl() (string, error) {
	absUrl := l.makeAbsURL()
	if _, err := url.Parse(absUrl); err != nil {
		return "", err
	}
	return absUrl, nil
}
