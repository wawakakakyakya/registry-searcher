package subCmd

import (
	"context"
	"flag"
	"registry-searcher/requester"

	"github.com/google/subcommands"
)

type ListImageCmd struct {
	*baseSubCmd
}

var listURL requester.URL

func (*ListImageCmd) Name() string     { return "list-image" }
func (*ListImageCmd) Synopsis() string { return "get image list" }

//flagはコマンドごとに異なる
func (l *ListImageCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&listURL.Address, "address", "", "target address example.com")
	f.IntVar(&listURL.Port, "port", 5000, "port number")
	f.IntVar(&listURL.Version, "version", 2, "registry version, default 2")
	f.BoolVar(&listURL.UseSsl, "usessl", true, "http(false) or https(default)")
}

func (l *ListImageCmd) Execute(mainCtx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	// 依存関係はここでまとめる
	lisRequester := requester.NewRequester(&listURL, "list")
	return l._Execute(mainCtx, lisRequester)
}
