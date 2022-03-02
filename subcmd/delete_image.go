package subcmd

import (
	"context"
	"flag"
	"registry-searcher/registry_url"
	"registry-searcher/requester"

	"github.com/google/subcommands"
)

type DeleteImageCmd struct {
	*baseSubCmd
}

var deleteImageUrl registry_url.DeleteImageUrl

func (*DeleteImageCmd) Name() string     { return "delete-image" }
func (*DeleteImageCmd) Synopsis() string { return "delete container image" }

//flagはコマンドごとに異なる
func (r *DeleteImageCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&deleteImageUrl.Address, "address", "", "target address example.com")
	f.IntVar(&deleteImageUrl.Port, "port", 5000, "port number")
	f.IntVar(&deleteImageUrl.Version, "version", 2, "registry version, default 2")
	f.BoolVar(&deleteImageUrl.UseSsl, "usessl", true, "http(false) or https(default)")
	f.StringVar(&deleteImageUrl.Image, "image", "", "specity image and tag(default: latest) ex: image:tag")
}

func (r *DeleteImageCmd) Execute(mainCtx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	// 依存関係はここでまとめる
	getManifestRequester := requester.NewDeleteImageRequester(&deleteImageUrl)
	return r._Execute(mainCtx, getManifestRequester)
}
