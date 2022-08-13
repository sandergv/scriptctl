package types

type Action struct {
	ID     string       `json:"id"`
	Name   string       `json:"name"`
	Script ScriptLookup `json:"script"`
	ExecID string       `json:"exec_id"`
}

type CreateActionRequest struct {
	Name     string `json:"name"`
	ScriptID string `json:"script_id"`
	ExecID   string `json:"exec_id"`
}

type CreateActionResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
	ID     string `json:"id"`
}

type GetActionResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
	Data   Action `json:"data"`
}

type GetActionListResponse struct {
	Status string   `json:"status"`
	Error  string   `json:"error"`
	Data   []Action `json:"data"`
}
