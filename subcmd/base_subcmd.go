package subcmd

import (
	"context"
	"fmt"
	"log"
	"registry-searcher/requester"

	"github.com/google/subcommands"
)

type baseSubCmd struct {
}

var resChan chan string
var errChan chan error

func (*baseSubCmd) Usage() string {
	return `print [-capitalize] <some text>:
  Print args to stdout.
`
}

//run in goroutine
func (l *baseSubCmd) execute(requester requester.RequesterInterface) {
	url, err := requester.GetUrl()
	if err != nil {
		log.Fatal("parse url failed")
		errChan <- err
		return
	}
	if res, err := requester.Execute(url); err != nil {
		errChan <- err
		return
	} else {
		resChan <- res
		return
	}
}

func (l *baseSubCmd) _Execute(mainCtx context.Context, requester requester.RequesterInterface) subcommands.ExitStatus {
	resChan = make(chan string)
	defer close(resChan)
	errChan = make(chan error)
	defer close(errChan)

	go l.execute(requester)

	select {
	case <-mainCtx.Done(): //cancel by timeout
		log.Println("context was timeout!")
		return subcommands.ExitFailure
	case res := <-resChan: // normal end
		fmt.Print(res)
		return subcommands.ExitSuccess
	case err := <-errChan: // err in execute
		fmt.Println(err.Error())
		return subcommands.ExitFailure
	}
}
