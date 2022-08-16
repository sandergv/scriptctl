package cli

import (
	"context"
	"fmt"

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
	Auth *AuthCMD `arg:"subcommand:auth" help:"Manage authentication"`

	// management commands
	Workspace *WorkspaceCMD `arg:"subcommand:workspace"`
	Script    *ScriptCMD    `arg:"subcommand:script" help:"Manage scripts"`
	Exec      *ExecCMD      `arg:"subcommand:exec" help:"Manage execution configs"`
	Namespace *NamespaceCMD `arg:"subcommand:namespace" help:"Manage namespaces"`
	Endpoint  *EndpointCMD  `arg:"subcommand:endpoint" help:"Manage endpoints"`
	Action    *ActionCMD    `arg:"subcommand:action" help:"Manage actions"`

	// comands
	Create  *CreateCMD  `arg:"subcommand:create" help:"Create resources with a configuration file"`
	Run     *RunCMD     `arg:"subcommand:run" help:"Run a script with the given options"`
	Version *VersionCMD `arg:"subcommand:version" help:"Show scriptlab version"`
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
	ctx = context.WithValue(ctx, ParserContextKey, p)

	switch {
	case args.Workspace != nil:
		args.Workspace.handle(ctx)

	case args.Script != nil:
		err := args.Script.handle(ctx)
		if err != nil {
			fmt.Println(err)
			p.FailSubcommand(err.Error(), "script")
		}
	case args.Exec != nil:
		err := args.Exec.handle(ctx)
		if err != nil {
			p.FailSubcommand(err.Error(), "exec")
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
	case args.Action != nil:
		err := args.Action.handle(ctx)
		if err != nil {
			p.FailSubcommand(err.Error(), "action")
		}
	case args.Run != nil:
		args.Run.handle(client)

	case args.Version != nil:
		args.Version.handle()
	}
}
