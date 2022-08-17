package cli

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/sandergv/scriptlab/pkg/scriptlabctl/types"
)

type ScriptCMD struct {
	Create *CreateScriptCMD `arg:"subcommand:create"`
	Update *UpdateScriptCMD `arg:"subcommand:update"`
	List   *ListScripCMD    `arg:"subcommand:list"`
}

func (s *ScriptCMD) handle(ctx context.Context) error {

	switch {
	case s.Create != nil:
		s.Create.handle(ctx)
	case s.Update != nil:
		s.Update.handle(ctx)

	case s.List != nil:
		s.List.handle(ctx)
	}

	return nil
}

type CreateScriptCMD struct {
	Name        string `arg:""`
	Description string `arg:"" help:"prueba de help"`
	Type        string `arg:""`
	FilePath    string `arg:"positional" placeholder:"FILE"`
}

func (c *CreateScriptCMD) handle(ctx context.Context) error {

	client := getClientFromContext(ctx)

	if c.FilePath == "" {
		return errors.New("file parameter is required")
	}

	content, err := os.ReadFile(c.FilePath)
	if err != nil {
		return err
	}

	fileName := filepath.Base(c.FilePath)

	id, err := client.CreateScript(types.CreateScriptOptions{
		Name:        c.Name,
		Description: c.Description,
		Type:        c.Type,
		FileName:    fileName,
		FileContent: string(content),
	})
	if err != nil {
		return err
	}
	fmt.Println("Script ID:", id)
	return nil
}

type UpdateScriptCMD struct {
	ID       string `arg:"positional,required"`
	FilePath string `arg:"positional,required" placeholder:"FILE"`
}

func (u *UpdateScriptCMD) handle(ctx context.Context) error {
	client := getClientFromContext(ctx)

	content, err := os.ReadFile(u.FilePath)
	if err != nil {
		return err
	}

	updatedAt, err := client.UpdateScript(types.UpdateScriptFileRequest{
		ID:          u.ID,
		FileContent: string(content),
	})
	if err != nil {
		return err
	}
	fmt.Println("Updated At:", updatedAt.Format(time.RFC3339Nano))
	return nil
}

type ListScripCMD struct {
}

func (l *ListScripCMD) handle(ctx context.Context) error {

	client := getClientFromContext(ctx)

	head := []string{"ID", "NAME", "DESCRIPTION", "TYPE", "CREATED", "UPDATED"}

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
		data = append(data, []string{v.ID, v.Name, v.Description, v.Type, time.Since(v.CreatedAt).Round(time.Second).String(), time.Since(v.UpdatedAt).Round(time.Second).String()})
	}

	showTable(head, data)

	return nil
}
