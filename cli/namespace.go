package cli

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"

	"github.com/sandergv/scriptlab/pkg/scriptlabctl"
	"github.com/sandergv/scriptlab/pkg/scriptlabctl/types"
)

type NamespaceCMD struct {
	Create  *CreateNamespaceCMD  `arg:"subcommand:create"`
	List    *ListNamespaceCMD    `arg:"subcommand:list"`
	Inspect *InspectNamespaceCMD `arg:"subcommand:inspect"`
}

func (n *NamespaceCMD) handle(ctx context.Context) error {

	switch {
	case n.Create != nil:
		err := n.Create.handle(ctx)
		if err != nil {
			return err
		}
	case n.List != nil:
		err := n.List.handle(ctx)
		if err != nil {
			return err
		}
	case n.Inspect != nil:
		err := n.Inspect.handle(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

type CreateNamespaceCMD struct {
	Name string   `arg:"positional"`
	Env  []string `arg:"-e"`
}

func (c *CreateNamespaceCMD) handle(ctx context.Context) error {

	if c.Name == "" {
		return errors.New("NAME parameter is required")
	}
	valid, _ := regexp.MatchString("^[a-z-_]{3,12}", c.Name)
	if !valid {
		return errors.New("invalid namespace name")
	}

	client := ctx.Value(ClientContextKey).(*scriptlabctl.Client)

	id, err := client.CreateNamespace(types.CreateNamespaceOptions{
		Name: c.Name,
		Env:  c.Env,
	})
	if err != nil {
		return err
	}
	fmt.Println("Namespace ID:", id)
	return nil
}

type ListNamespaceCMD struct {
}

func (l *ListNamespaceCMD) handle(ctx context.Context) error {

	client := getClientFromContext(ctx)

	nss, err := client.GetNamespaceList()
	if err != nil {
		return err
	}

	header := []string{"ID", "NAME", "CONTEXT"}

	data := [][]string{}
	for _, v := range nss {
		data = append(data, []string{v.ID, v.Name, v.Context.Name})
	}

	showTable(header, data)

	return nil
}

type InspectNamespaceCMD struct {
	Name string `arg:"positional,required"`
}

func (i *InspectNamespaceCMD) handle(ctx context.Context) error {

	client := getClientFromContext(ctx)

	ns, err := client.GetNamespace(i.Name)
	if err != nil {
		return err
	}

	b, _ := json.MarshalIndent(ns, "", "  ")

	fmt.Println(string(b))

	return nil
}
