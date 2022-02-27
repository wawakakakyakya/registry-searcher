package requester

type RequesterInterface interface {
	Get(url string, resChan chan string) error
	GetUrl() (string, error)
}
