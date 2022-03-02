package registry_url

import (
	"fmt"
	"strings"
)

type DeleteImageUrl struct {
	Address string
	Port    int
	Version int
	UseSsl  bool
	Digest  string
	Image   string
	*baseRegistryURL
}

func (r *DeleteImageUrl) makeAbsURL() string {
	baseUrl := fmt.Sprintf("%s://%s", r.getScheme(r.UseSsl), r.getAuthority(r.Address, r.Port))
	image := strings.Split(r.Image, ":")[0]
	return fmt.Sprintf("%s/%s/%s/manifests/%s", baseUrl, r.getVersion(r.Version), image, r.Digest)
}

//entry point
func (r *DeleteImageUrl) MakeUrl() (string, error) {
	absUrl := r.makeAbsURL()
	return r.parse(absUrl)
}
