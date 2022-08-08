package types

type EnvLookup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Exec struct {
	ID string `json:"id"`
	// Name    string    `json:"name"`
	ExecEnv EnvLookup `json:"exec_env"`
	Type    string    `json:"type"`
	File    string    `json:"file"`
	Env     []string  `json:"env"`
	Args    []string  `json:"args"`
}

type CreateExecRequest struct {
	// Name        string   `json:"name"`
	ExecEnv  string   `json:"exec_env,omitempty"` // can be id or name, both are unique
	ScriptID string   `json:"script_id"`
	Context  string   `json:"context"`
	Env      []string `json:"env"`
	Args     []string `json:"args"`
}

type CreateExecResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
	ID     string `json:"id"`
}

type GetExecResponse struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	Data   Exec   `json:"data"`
}

type GetExecListResponse struct {
	Status string `json:"status"`
	Data   []Exec `json:"data"`
}

type RemoveExecResponse struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}
