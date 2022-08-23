package types

type Action struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Namespace   NamespaceLookup `json:"namespace"`
	Script      ScriptLookup    `json:"script"`
}

type CreateActionRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Namespace   string `json:"namespace"`
	ScriptID    string `json:"script_id"`
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
