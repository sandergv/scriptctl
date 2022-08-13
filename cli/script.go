package cli

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"sort"
	"time"

	"github.com/sandergv/scriptlab/pkg/scriptlabctl/types"
)

type ScriptCMD struct {
	Create *CreateScriptCMD `arg:"subcommand:create"`
	List   *ListScripCMD    `arg:"subcommand:list"`
}

func (s *ScriptCMD) handle(ctx context.Context) error {

	if s.Create != nil {
		s.Create.handle(ctx)
	}
	switch {
	case s.Create != nil:
		s.Create.handle(ctx)
	case s.List != nil:
		s.List.handle(ctx)
	}

	return nil
}

type CreateScriptCMD struct {
	Name        string `arg:""`
	Description string `arg:""`
	Type        string `arg:""`
	FilePath    string `arg:"positional" placeholder:"FILE"`
}

func (c *CreateScriptCMD) handle(ctx context.Context) error {

	client := getClientFromContext(ctx)

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

type ListScripCMD struct {
}

func (l *ListScripCMD) handle(ctx context.Context) error {

	client := getClientFromContext(ctx)

	head := []string{"ID", "NAME", "TYPE", "CREATED"}

	data := [][]string{}

	scripts, err := client.GetScriptList()
	if err != nil {
		fmt.Println(err)
		return err
	}

	sort.Slice(scripts, func(i, j int) bool {
		return scripts[i].CreatedAt.After(scripts[j].CreatedAt)
	})

	for _, v := range scripts {
		data = append(data, []string{v.ID, v.Name, v.Type, time.Since(v.CreatedAt).Round(time.Second).String()})
	}

	showTable(head, data)

	return nil
}
