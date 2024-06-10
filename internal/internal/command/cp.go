package command

import (
	"context"

	"github.com/goexl/args"
	"github.com/goexl/gex"
	"github.com/goexl/gox"
	"github.com/goexl/log"
)

type Copy struct {
	logger log.Logger
	_      gox.CannotCopy
}

func NewCopy(logger log.Logger) *Copy {
	return &Copy{
		logger: logger,
	}
}

func (c *Copy) Exec(ctx context.Context, source string, target string) (err error) {
	arguments := args.New().Build().Flag("recursive").Subcommand(source).Subcommand(target)
	_, err = gex.New("cp").Context(ctx).Arguments(arguments.Build()).Build().Exec()

	return
}
