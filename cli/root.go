package cli

import (
	"github.com/alexflint/go-arg"
	"github.com/sandergv/scriptctl/pkg/scriptlabctl"
)

type Globals struct {
}

type args struct {
	Auth *AuthCMD `arg:"subcommand:auth"`

	Run *RunCMD `arg:"subcommand:run"`
}

func (args) Version() string {
	return "scriptctl 0.0.1"
}

func Exec(client *scriptlabctl.Client) {
	var args args
	arg.MustParse(&args)
	switch {
	case args.Auth != nil:
		args.Auth.handle(client)
	case args.Run != nil:
		args.Run.handle(client)
	}
}
