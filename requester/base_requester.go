package requester

import (
	"errors"
	"fmt"
	"net/http"
)

type baseRequester struct {
}

func (r *baseRequester) get(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ok := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !ok {
		fmt.Printf("status code(%d) was not succeeded\n", resp.StatusCode)
		return "", errors.New(printBody(resp.Body))
	}

	return printBody(resp.Body), nil
}
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("status code(%d) was not succeeded\n", resp.StatusCode)
		return errors.New(printBody(resp.Body))
	}

	resChan <- printBody(resp.Body)
	return nil
}
