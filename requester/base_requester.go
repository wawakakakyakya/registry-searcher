package requester

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

type URL struct {
	Address string
	Port    int
	Version int
	UseSsl  bool
}

type baseRequester struct {
}

func (r *baseRequester) getVersion(url *URL) string {
	return fmt.Sprintf("v%d", url.Version)
}

func (r *baseRequester) scheme(url *URL) string {
	scheme := "http"
	if url.UseSsl {
		scheme = "https"
	}
	return scheme
}

func (r *baseRequester) authority(url *URL) string {
	return fmt.Sprintf("%s:%d", url.Address, url.Port)
}

func (r *baseRequester) _MakeURL(url *URL) (string, error) {
	baseUrl := fmt.Sprintf("%s://%s", r.scheme(url), r.authority(url))
	u, err := r.parse(baseUrl, r.getVersion(url))
	if err != nil {
		fmt.Println("url parse error")
		return "", err
	}
	return u, nil
}

func (r *baseRequester) parse(base string, version string) (string, error) {
	u, err := url.Parse(base)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, version, "_catalog")
	return u.String(), nil
}

func (r *baseRequester) Get(url string, resChan chan string) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("status code(%d) was not succeeded\n", resp.StatusCode)
		return errors.New(printBody(resp.Body))
	}

	resChan <- printBody(resp.Body)
	return nil
}

func NewRequester(url *URL, spec string) RequesterInterface {
	switch spec {
	case "list":
		return &ListImageRequester{url: url}
	default:
		return nil
	}
}
