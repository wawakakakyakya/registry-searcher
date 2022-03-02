package requester

type RequesterInterface interface {
	GetUrl() (string, error)
	Execute(url string) (string, error)
}
