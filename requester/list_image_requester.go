package requester

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"path"
)

func printBody(body io.ReadCloser) string {
	_body, _ := ioutil.ReadAll(body)
	// body は []byte
	//fmt.Println(_body)
	return string(_body)
}

type ListImageRequester struct {
	// url URL
	url *URL
	*baseRequester
}

func (l *ListImageRequester) MakeURL() (string, error) {
	url, err := l._MakeURL(l.url)
	if err != nil {
		fmt.Println(err.Error())
	}
	return url, nil
}

func (l *ListImageRequester) parse(base string, version string) (string, error) {
	u, err := url.Parse(base)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, version, "_catalog")
	return u.String(), nil
}

func NewListImageRequester(url *URL) RequesterInterface {
	return &ListImageRequester{url: url}
}
