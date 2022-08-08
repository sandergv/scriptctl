package cli

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/sandergv/scriptctl/pkg/scriptlabctl"
	"github.com/sandergv/scriptctl/pkg/scriptlabctl/types"
)

type CreateScriptCMD struct {
	Name        string `arg:""`
	Description string `arg:""`
	Type        string `arg:""`
	FilePath    string `arg:"positional" placeholder:"FILE"`
}

func (c *CreateScriptCMD) handle(ctx context.Context) error {

	client := ctx.Value(ClientContextKey).(*scriptlabctl.Client)

	if c.FilePath == "" {
		return errors.New("file parameter is required")
	}

	fileName := filepath.Base(c.FilePath)

	id, err := client.CreateScript(types.CreateScriptOptions{
		Name:     c.Name,
		Type:     c.Type,
		FileName: fileName,
	})
	if err != nil {
		return err
	}
	fmt.Println(id)
	return nil
}
