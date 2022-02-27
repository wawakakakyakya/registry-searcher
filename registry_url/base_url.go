package registry_url

import (
	"fmt"
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
