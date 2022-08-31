package cli

import (
	"context"
	"fmt"

	"github.com/sandergv/scriptlab/pkg/scriptlabctl/types"
)

type EndpointCMD struct {
	Create *CreateEndpointCMD `arg:"subcommand:create"`
	List   *ListEndpointCMD   `arg:"subcommand:list"`
}

func (e *EndpointCMD) handle(ctx context.Context) error {

	p := getParserFromContext(ctx)

	switch {
	case e.Create != nil:
		err := e.Create.handle(ctx)
		if err != nil {
			p.FailSubcommand(err.Error(), "endpoint", "create")
		}

	case e.List != nil:
		err := e.List.handle(ctx)
		if err != nil {
			p.FailSubcommand(err.Error(), "endpoint", "list")
		}
	}

	return nil
}

type CreateEndpointCMD struct {
	Name      string `arg:"positional,required"`
	Api       string `arg:"" default:"private"`
	Method    string `arg:"" default:"get"`
	Namespace string `arg:"-n"`
	Script    string `arg:"positional,required"`
}

func (c *CreateEndpointCMD) handle(ctx context.Context) error {

	client := getClientFromContext(ctx)

	private := true
	if c.Api == "public" {
		private = false
	}

	id, err := client.CreateEndpoint(types.CreateEndpointOptions{
		Name:      c.Name,
		Namespace: c.Namespace,
		Method:    c.Method,
		Private:   private,
		ScriptID:  c.Script,
	})
	if err != nil {
		return err
	}
	fmt.Println("Endpoint ID", id)

	return nil
}

type ListEndpointCMD struct {
	Namespace string `arg:"-n,--namespace"`
}

func (l *ListEndpointCMD) handle(ctx context.Context) error {

	client := getClientFromContext(ctx)

	eps, err := client.GetEndpointList(l.Namespace)
	if err != nil {
		return err
	}

	// order namespace's endpoints
	ns := map[string][]types.Endpoint{}
	for _, e := range eps {
		if v, ok := ns[e.Namespace.Name]; ok {
			v = append(v, e)
			ns[e.Namespace.Name] = v
		} else {
			ns[e.Namespace.Name] = []types.Endpoint{
				e,
			}
		}
	}

	header := []string{"ID", "NAME", "METHOD", "API", "NAMESPACE", "SCRIPT"}

	data := [][]string{}

	for _, n := range ns {
		for _, e := range n {
			api := ""
			if e.Private {
				api = "private"
			} else {
				api = "public"
			}
			data = append(data,
				[]string{e.ID, e.Name, e.Method, api, e.Namespace.Name, e.Script.Name},
			)
		}
	}

	showTable(header, data)

	return nil
}
