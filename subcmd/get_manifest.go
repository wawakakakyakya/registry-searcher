package subcmd

import (
	"context"
	"flag"
	"registry-searcher/registry_url"
	"registry-searcher/requester"

	"github.com/google/subcommands"
)

type GetManifestCmd struct {
	*baseSubCmd
}

var getManifestUrl registry_url.GetManifestUrl

func (*GetManifestCmd) Name() string     { return "manifest" }
func (*GetManifestCmd) Synopsis() string { return "get container manifest" }

//flagはコマンドごとに異なる
func (l *GetManifestCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&getManifestUrl.Address, "address", "", "target address example.com")
	f.IntVar(&getManifestUrl.Port, "port", 5000, "port number")
	f.IntVar(&getManifestUrl.Version, "version", 2, "registry version, default 2")
	f.BoolVar(&getManifestUrl.UseSsl, "usessl", true, "http(false) or https(default)")
	f.StringVar(&getManifestUrl.Tag, "tag", "latest", "image tag(default: latest)")
	f.StringVar(&getManifestUrl.Image, "image", "", "image name")

}

func (l *GetManifestCmd) Execute(mainCtx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	// 依存関係はここでまとめる

	getManifestRequester := requester.NewGetManifestRequester(&getManifestUrl)
	return l._Execute(mainCtx, getManifestRequester)
}
