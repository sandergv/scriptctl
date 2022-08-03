package types

type RunCodeOptions struct {
	ExecEnv string
	Type    string
	Env     []string
	Args    []string
	Body    map[string]interface{}
	Code    string
}

type RunExecOptions struct {
	ExecID string
	Body   map[string]interface{}
}

type RunExecRequest struct {
	ID   string                 `json:"id"`
	Body map[string]interface{} `json:"body"`
}

type RunCodeRequest struct {

	// required
	ExecEnv string `json:"exec_env"`

	// required
	Type string `json:"type"`

	// optional
	Envs []string `json:"envs"`

	// optional
	Args []string `json:"args"`

	// required
	Content string `json:"content"`
}

type RunDetails struct {
	ExitCode int      `json:"exit_code"`
	Error    string   `json:"error,omitempty"`
	Output   []string `json:"output"`
	Logs     []string `json:"logs"`
	Response string   `json:"response"`
}

type RunResponse struct {
	Status  string     `json:"status"`
	Error   string     `json:"error"`
	Details RunDetails `json:"details"`
}
