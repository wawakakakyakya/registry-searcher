package registry_url

import (
	"fmt"
	"net/url"
)

type baseRegistryURL struct {
}

func (b *baseRegistryURL) getVersion(version int) string {
	return fmt.Sprintf("v%d", version)
}

func (b *baseRegistryURL) getScheme(useSsl bool) string {
	scheme := "http"
	if useSsl {
		scheme = "https"
	}
	return scheme
}

func (b *baseRegistryURL) getAuthority(addr string, port int) string {
	return fmt.Sprintf("%s:%d", addr, port)
}

func (b *baseRegistryURL) parse(absUrl string) (string, error) {
	if _, err := url.Parse(absUrl); err != nil {
		return "", err
	}
	return absUrl, nil
}
