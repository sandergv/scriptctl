package cli

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/sandergv/scriptlab/pkg/scriptlabctl"
	"github.com/sandergv/scriptlab/pkg/scriptlabctl/types"
)

type NamespaceCMD struct {
	Create *CreateNamespaceCMD `arg:"subcommand:create"`
	List   *ListNamespaceCMD   `arg:"subcommand:list"`
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
	}
	return nil
}

type CreateNamespaceCMD struct {
	Name string `arg:"positional"`
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

	header := []string{"ID", "NAME", "ENDPOINTS"}

	data := [][]string{}
	for _, v := range nss {
		data = append(data, []string{v.ID, v.Name, strconv.Itoa(len(v.Endpoints))})
	}

	showTable(header, data)

	return nil
}
