package types

type Command struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Namespace   string       `json:"namespace"`
	Env         []string     `json:"env"`
	Script      ScriptLookup `json:"script"`
	Context     string       `json:"context"`
}

type CreateCommandRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Namespace   string   `json:"namespace"`
	Env         []string `json:"env"`
	ScriptID    string   `json:"script_id"`
}

type CreateCommandResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
	ID     string `json:"id"`
}

type GetCommandResponse struct {
	Status string  `json:"status"`
	Error  string  `json:"error"`
	Data   Command `json:"data"`
}

type GetCommandListResponse struct {
	Status string    `json:"status"`
	Error  string    `json:"error"`
	Data   []Command `json:"data"`
}

type RemoveCommandResponse struct {
	Status string `json:"status"`
}

type RunCommandRequest struct {
	Name string   `json:"name"`
	Args []string `json:"args"`
}

type RunCommandResponse struct {
	Status  string     `json:"status"`
	Error   string     `json:"error"`
	Details RunDetails `json:"details"`
}
