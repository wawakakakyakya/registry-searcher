package requester

type RequesterInterface interface {
	getVersion(url *URL) string
	scheme(url *URL) string
	authority(url *URL) string
	MakeURL() (string, error)
	parse(base string, version string) (string, error)
	Get(url string, resChan chan string) error
}
