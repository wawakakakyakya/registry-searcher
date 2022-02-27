package requester

import (
	"errors"
	"fmt"
	"net/http"
)

type baseRequester struct {
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
