package cli

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sandergv/scriptctl/pkg/scriptlabctl"
	"github.com/sandergv/scriptctl/pkg/scriptlabctl/types"
)

type CreateCMD struct {
	Script *CreateScriptCMD `arg:"subcommand:script"`

	FilePath string `arg:"--file"`
}

func (c *CreateCMD) handle(ctx context.Context) error {

	if c.FilePath == "" {
		return errors.New("FILEPATH parameter is required if no command is provide")
	}

	fileInfo, err := os.Stat(c.FilePath)
	if os.IsNotExist(err) {
		return err
	}
	if fileInfo.IsDir() {
		return errors.New("the path is a directory, not a file")
	}

	return nil
}

func create(ctx context.Context, cfg ConfigFile) error {

	client := ctx.Value(ClientContextKey).(*scriptlabctl.Client)

	scriptId := ""

	if cfg.Script != nil {

		f, err := os.Stat(cfg.Script.FilePath)
		if os.IsNotExist(err) {
			return errors.New("file does not exist")
		}
		if f.IsDir() {
			return errors.New("the path is a directory")
		}

		fileName := filepath.Base(cfg.Script.FilePath)

		content, err := ioutil.ReadFile(cfg.Script.FilePath)
		if err != nil {
			return err
		}
		scriptId, err = client.CreateScript(types.CreateScriptOptions{
			Name: cfg.Script.Name,
			// description
			Type:        cfg.Script.Type,
			FileName:    fileName,
			FileContent: string(content),
		})
		fmt.Println("Script ID:", scriptId)
	}

	return nil
}
