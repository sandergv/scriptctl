package cli

//
type ScriptConfig struct {
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description" json:"description"`
	// Required
	Type string `yaml:"type" json:"type"`

	// Required
	FilePath string `yaml:"file-path" json:"file_path"`

	// Optional
	DefaultExec bool `yaml:"default-exec" json:"default_exec"`
}

type ExecConfig struct {

	//ExecEnv
	//
	// Optional, if is empty default env will be assign
	ExecEnv string `yaml:"exec-env" json:"exec_env"`

	// Optional
	Env map[string]interface{} `yaml:"env" json:"env"`

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

type ConfigFile struct {
	Script  *ScriptConfig  `yaml:"script" json:"script"`
	Exec    *ExecConfig    `yaml:"exec" json:"exec"`
	Context *ContextConfig `yaml:"context" json:"context"`
	// Namespace
	Endpoint *EndpointConfig `yaml:"endpoint" json:"endpoint"`
}
