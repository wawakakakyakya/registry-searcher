package subcmd

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

type subCmdInterface interface {
	Name() string
	Synopsis() string
	Usage() string
	SetFlags(f *flag.FlagSet)
	execute()
	Execute(mainCtx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus
}
