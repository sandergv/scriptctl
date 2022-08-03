package cli

//
type ScriptConfig struct {

	// Required
	Type string `yaml:"type" json:"type"`

	// Required
	FileName string `yaml:"file-name" json:"file_name"`

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

type ConfigFile struct {
	Script  **ScriptConfig `yaml:"script" json:"script"`
	Exec    *ExecConfig    `yaml:"exec" json:"exec"`
	Context *ContextConfig `yaml:"context" json:"context"`
}
