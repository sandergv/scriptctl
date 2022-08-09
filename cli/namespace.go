package cli

import (
	"context"
	"errors"
	"fmt"

	"github.com/sandergv/scriptctl/pkg/scriptlabctl"
	"github.com/sandergv/scriptctl/pkg/scriptlabctl/types"
)

type CreateNamespaceCMD struct {
	Name string `arg:"positional"`
}

func (c *CreateNamespaceCMD) handle(ctx context.Context) error {

	if c.Name == "" {
		return errors.New("NAME parameter is required")
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
