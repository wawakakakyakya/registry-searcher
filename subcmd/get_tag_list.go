package subcmd

import (
	"context"
	"flag"
	"registry-searcher/registry_url"
	"registry-searcher/requester"

	"github.com/google/subcommands"
)

type GetTagCmd struct {
	*baseSubCmd
}

var getTagUrl registry_url.GetTagtUrl

func (*GetTagCmd) Name() string     { return "get-tag" }
func (*GetTagCmd) Synopsis() string { return "get image taglist" }

//flagはコマンドごとに異なる
func (l *GetTagCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&getTagUrl.Address, "address", "", "target address example.com")
	f.IntVar(&getTagUrl.Port, "port", 5000, "port number")
	f.IntVar(&getTagUrl.Version, "version", 2, "registry version, default 2")
	f.BoolVar(&getTagUrl.UseSsl, "usessl", true, "http(false) or https(default)")
	f.StringVar(&getTagUrl.Image, "image", "", "image name")
}

func (l *GetTagCmd) Execute(mainCtx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	// 依存関係はここでまとめる
	getTagRequester := requester.NewGetTagRequester(&getTagUrl)
	return l._Execute(mainCtx, getTagRequester)
}
