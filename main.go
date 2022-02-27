package main

import (
	"context"
	"flag"
	"os"
	"registry-searcher/subcmd"
	"time"

	"github.com/google/subcommands"
)

func init() {
	subcommands.Register(&subcmd.ListImageCmd{}, "List command register")
	subcommands.Register(&subcmd.GetManifestCmd{}, "Get container manifest")
	subcommands.Register(&subcmd.GetTagCmd{}, "Get tag list")
	flag.Parse()
}

func main() {

	mainCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	eCode := int(subcommands.Execute(mainCtx))

	os.Exit(eCode)

}
