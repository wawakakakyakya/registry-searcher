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

// func (r *Requester) getVersion() string {
// 	return fmt.Sprintf("v%d", r.Version)
// }

// func (r *Requester) scheme() string {
// 	scheme := "http"
// 	if r.UseSsl {
// 		scheme = "https"
// 	}
// 	return scheme
// }

// func (r *Requester) authority() string {
// 	return fmt.Sprintf("%s:%d", r.Address, r.Port)
// }

func (l *ListImageRequester) MakeURL() (string, error) {
	url, err := l._MakeURL(l.url)
	if err != nil {
		fmt.Println(err.Error())
	}
	return url, nil
	// baseUrl := fmt.Sprintf("%s://%s", l.scheme(url), l.authority(url))
	// u, err := l.parse(baseUrl, l.getVersion(url))
	// if err != nil {
	// 	fmt.Println("url parse error")
	// 	return "", err
	// }
	// return u, nil
}

func (l *ListImageRequester) parse(base string, version string) (string, error) {
	u, err := url.Parse(base)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, version, "_catalog")
	return u.String(), nil
}

// func (r *Requester) Get(url string, resChan chan string) error {
// 	resp, err := http.Get(url)

// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != 200 {
// 		fmt.Printf("status code(%d) was not succeeded\n", resp.StatusCode)
// 		return errors.New(printBody(resp.Body))
// 	}

// 	resChan <- printBody(resp.Body)
// 	return nil
// }
