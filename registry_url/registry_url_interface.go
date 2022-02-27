package registry_url

type RegistryUrlInterface interface {
	getVersion(version int) string
	getScheme(useSsl bool) string
	getAuthority(addr string, port int) string
	makeAbsURL() string
	MakeUrl() (string, error)
}
