package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/alexflint/go-arg"
	"github.com/sandergv/scriptlab/pkg/scriptlabctl"
)

const (
	ClientContextKey = iota
	ParentCommandContextKey
	ParserContextKey
)

func getClientFromContext(ctx context.Context) *scriptlabctl.Client {
	return ctx.Value(ClientContextKey).(*scriptlabctl.Client)
}

func getParserFromContext(ctx context.Context) *arg.Parser {
	return ctx.Value(ParserContextKey).(*arg.Parser)
}

type args struct {
	Auth   *AuthCMD   `arg:"subcommand:auth"`
	Create *CreateCMD `arg:"subcommand:create"`

	Script    *ScriptCMD    `arg:"subcommand:script"`
	Namespace *NamespaceCMD `arg:"subcommand:namespace"`
	Endpoint  *EndpointCMD  `arg:"subcommand:endpoint"`

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
	ctx = context.WithValue(ctx, ParentCommandContextKey, p)

	switch {
	case args.Auth != nil:
		args.Auth.handle(client)

	case args.Script != nil:
		err := args.Script.handle(ctx)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error:", err.Error())
			p.WriteHelpForSubcommand(os.Stdout, "script")
		}

	case args.Create != nil:
		err := args.Create.handle(ctx)
		if err != nil {
			p.FailSubcommand(err.Error(), "create")
		}
	case args.Namespace != nil:
		err := args.Namespace.handle(ctx)
		if err != nil {
			p.FailSubcommand(err.Error(), "namespace")
		}
	case args.Endpoint != nil:
		err := args.Endpoint.handle(ctx)
		if err != nil {
			p.FailSubcommand(err.Error(), "endpoint")
		}

	case args.Run != nil:
		args.Run.handle(client)

	case args.Version != nil:
		args.Version.handle()
	}
}
