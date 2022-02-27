package subcmd

import (
	"context"
	"flag"
	"registry-searcher/registry_url"
	"registry-searcher/requester"

	"github.com/google/subcommands"
)

type ListImageCmd struct {
	*baseSubCmd
}

var listImageUrl registry_url.ListImageUrl

func (*ListImageCmd) Name() string     { return "list-image" }
func (*ListImageCmd) Synopsis() string { return "get image list" }

//flagはコマンドごとに異なる
func (l *ListImageCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&listImageUrl.Address, "address", "", "target address example.com")
	f.IntVar(&listImageUrl.Port, "port", 5000, "port number")
	f.IntVar(&listImageUrl.Version, "version", 2, "registry version, default 2")
	f.BoolVar(&listImageUrl.UseSsl, "usessl", true, "http(false) or https(default)")
}

func (l *ListImageCmd) Execute(mainCtx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	// 依存関係はここでまとめる

	lisRequester := requester.NewListImageRequester(&listImageUrl)
	return l._Execute(mainCtx, lisRequester)
}
