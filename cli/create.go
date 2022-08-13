// TODO: every error must return the command that has the error

package cli

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/sandergv/scriptlab/pkg/scriptlabctl"
	"github.com/sandergv/scriptlab/pkg/scriptlabctl/types"
)

type CreateCMD struct {
	// Script *CreateScriptCMD `arg:"subcommand:script"`

	// Namespace *CreateNamespaceCMD `arg:"subcommand:namespace"`

	FilePath string `arg:"positional"`
}

func (c *CreateCMD) handle(ctx context.Context) error {

	var err error

	if c.FilePath == "" {
		return errors.New("FILEPATH parameter is required if no command is provide")
	}
	createFromConfig(ctx, c.FilePath)

	return err
}

// createFromConfig creates create entities from a config file
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
		if scriptId != "" {
			cfg.Exec.Script = scriptId
		}
		execId, _ = createExec(client, *cfg.Exec)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println("Exec ID:", execId)
	}
	if cfg.Endpoint != nil {
		if execId != "" {
			cfg.Endpoint.Exec = execId
		}
		epId, err := createEndpoint(client, *cfg.Endpoint)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println("Endpoint ID:", epId)
	}

	return nil
}

func createScript(client *scriptlabctl.Client, cfg ScriptConfig) (string, error) {
	// fmt.Println(cfg)
	fileName := filepath.Base(cfg.FilePath)

	if cfg.Type == "" {
		return "", errors.New("type parameter can't be empty")
	}

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

func createExec(client *scriptlabctl.Client, cfg ExecConfig) (string, error) {

	return client.CreateExec(types.CreateExecRequest{
		ExecEnv:  cfg.ExecEnv,
		ScriptID: cfg.Script,
		Env:      cfg.Env,
		Args:     cfg.Args,
	})
}

func createEndpoint(client *scriptlabctl.Client, cfg EndpointConfig) (string, error) {
	if cfg.Name == "" {
		return "", errors.New("name parameter is required")
	}
	if cfg.Namespace == "" {
		return "", errors.New("namespace parameter is required")
	}
	if cfg.Method == "" {
		return "", errors.New("method parameter is required")
	}
	if cfg.Exec == "" {
		return "", errors.New("exec parameter is required")
	}

	// check if namespace exist
	nss, err := client.GetNamespaceList()
	if err != nil {
		return "", err
	}
	exist := false
	for _, n := range nss {
		if n.ID == cfg.Namespace || n.Name == cfg.Namespace {
			exist = true
			break
		}
	}

	if !exist {
		nId, err := client.CreateNamespace(types.CreateNamespaceOptions{
			Name: cfg.Namespace,
		})
		if err != nil {
			return "", err
		}
		fmt.Println("Namespace ID:", nId)
	}

	return client.CreateEndpoint(types.CreateEndpointOptions{
		Name:      cfg.Name,
		Namespace: cfg.Namespace,
		Private:   cfg.Private,
		Method:    cfg.Method,
		ExecID:    cfg.Exec,
	})
}
