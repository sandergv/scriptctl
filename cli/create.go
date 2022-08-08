package cli

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/sandergv/scriptctl/pkg/scriptlabctl"
	"github.com/sandergv/scriptctl/pkg/scriptlabctl/types"
)

type CreateCMD struct {
	Script *CreateScriptCMD `arg:"subcommand:script"`

	FilePath string `arg:"--file"`
}

func (c *CreateCMD) handle(ctx context.Context) error {

	switch {
	case c.Script != nil:
		// do something
	default:
		fmt.Println("aqui")
		if c.FilePath == "" {
			return errors.New("FILEPATH parameter is required if no command is provide")
		}
		createFromConfig(ctx, c.FilePath)
	}
	return nil
}

// createFromConfig creates create entities fgrom a config file
func createFromConfig(ctx context.Context, fp string) error {

	cfg, err := parseConfig(fp)
	if err != nil {
		return err
	}

	client := ctx.Value(ClientContextKey).(*scriptlabctl.Client)
	scriptId := ""
	execId := ""
	if cfg.Script != nil {
		scriptId, err = createScript(client, *cfg.Script)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println("Script ID:", scriptId)
	}
	if cfg.Exec != nil {
		execId, _ = createExec(client, scriptId, *cfg.Exec)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println("Exec ID:", execId)
	}

	return nil
}

func createScript(client *scriptlabctl.Client, cfg ScriptConfig) (string, error) {
	// fmt.Println(cfg)
	fileName := filepath.Base(cfg.FilePath)

	content, err := ioutil.ReadFile(cfg.FilePath)
	if err != nil {
		fmt.Println("aqu")
		return "", err
	}
	return client.CreateScript(types.CreateScriptOptions{
		Name: cfg.Name,
		// description
		Type:        cfg.Type,
		FileName:    fileName,
		FileContent: string(content),
	})
}

func createExec(client *scriptlabctl.Client, sid string, cfg ExecConfig) (string, error) {

	return client.CreateExec(types.CreateExecRequest{
		ExecEnv:  cfg.ExecEnv,
		ScriptID: sid,
		Env:      cfg.Env,
		Args:     cfg.Args,
	})
}
