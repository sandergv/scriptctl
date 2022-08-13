package cli

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

//
type ScriptConfig struct {
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description" json:"description"`
	// Required
	Type string `yaml:"type" json:"type"`

	// Required
	FilePath string `yaml:"path" json:"path"`

	// Optional
	DefaultExec bool `yaml:"default-exec" json:"default_exec"`
}

type ExecConfig struct {

	// Script
	//
	// Required
	Script string `yaml:"script" json:"script"`

	//ExecEnv
	//
	// Optional, if is empty default env will be assign
	ExecEnv string `yaml:"exec-env" json:"exec_env"`

	// Optional
	EnvFile string `yaml:"env-file" json:"env_file"`

	// Optional
	Env []string `yaml:"env" json:"env"`

	// Optional
	Args []string `yaml:"args" json:"args"`

	// Context
	//
	// Optional
	Context string `yaml:"context" json:"context"`
}

type ContextConfig struct {

	// Required
	Name string `yaml:"name" json:"name"`

	// Optional
	Data map[string]interface{} `yaml:"data" json:"data"`
}

type NamespaceConfig struct {
	Name string `yaml:"name" json:"name"`
}

type EndpointConfig struct {
	Name      string `yaml:"name" json:"name"`
	Namespace string `yaml:"namespace" json:"namespace"`
	Private   bool   `yaml:"private" json:"private"`
	Method    string `yaml:"method" json:"method"`
	Exec      string `yaml:"exec" json:"exec"`
}

type ActionConfig struct {
	Name string `yaml:"name" json:"name"`
	Exec string `yaml:"exec" json:"exec"`
}

type ConfigFile struct {
	Script  *ScriptConfig  `yaml:"script" json:"script"`
	Exec    *ExecConfig    `yaml:"exec" json:"exec"`
	Context *ContextConfig `yaml:"context" json:"context"`
	// Namespace
	Endpoint *EndpointConfig `yaml:"endpoint" json:"endpoint"`
	Action   *ActionConfig   `yaml:"action" json:"action"`
}

func parseConfig(fp string) (ConfigFile, error) {

	fileInfo, err := os.Stat(fp)
	if os.IsNotExist(err) {
		fmt.Println("aqui")
		return ConfigFile{}, err
	}
	if fileInfo.IsDir() {
		return ConfigFile{}, errors.New("the path is a directory, not a file")
	}

	file, err := ioutil.ReadFile(fp)
	if err != nil {
		fmt.Println("aqui")
		return ConfigFile{}, err
	}

	cfg := ConfigFile{}

	err = yaml.Unmarshal(file, &cfg)

	return cfg, err
}

func handleEnvFile(fp string) ([]string, error) {

	fileInfo, err := os.Stat(fp)
	if os.IsNotExist(err) {
		fmt.Println("aqui")
		return nil, err
	}
	if fileInfo.IsDir() {
		return nil, errors.New("the path is a directory, not a file")
	}

	content, err := ioutil.ReadFile(fp)
	if err != nil {
		return nil, err
	}

	envList := strings.Split(string(content), "\n")
	return envList, nil
}
