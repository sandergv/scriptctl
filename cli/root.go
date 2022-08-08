package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/alexflint/go-arg"
	"github.com/sandergv/scriptctl/pkg/scriptlabctl"
)

const (
	ClientContextKey = iota
	ParentCommandContextKey
)

type Globals struct {
}

type args struct {
	Auth   *AuthCMD   `arg:"subcommand:auth"`
	Create *CreateCMD `arg:"subcommand:create"`

	Run     *RunCMD     `arg:"subcommand:run"`
	Version *VersionCMD `arg:"subcommand:version"`
}

// func (args) Version() string {
// 	return "scriptctl 0.0.1"
// }

func Exec(client *scriptlabctl.Client) {
	var args args
	p := arg.MustParse(&args)

	ctx := context.Background()
	ctx = context.WithValue(ctx, ClientContextKey, client)
	ctx = context.WithValue(ctx, ParentCommandContextKey, "root")

	switch {
	case args.Auth != nil:
		args.Auth.handle(client)

	case args.Create != nil:
		err := args.Create.handle(ctx)
		if err != nil {
			fmt.Println("Error:", err.Error())
			p.WriteHelpForSubcommand(os.Stdout, "create")
		}

	case args.Run != nil:
		args.Run.handle(client)

	case args.Version != nil:
		args.Version.handle()
	}
}
