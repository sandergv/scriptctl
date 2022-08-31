package types

type NamespaceLookup struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Endpoint struct {
	ID        string          `json:"id"`
	Name      string          `json:"name"`
	Namespace NamespaceLookup `json:"namespace,omitempty"`
	Private   bool            `json:"private"`
	Method    string          `json:"method"`
	Script    ScriptLookup    `json:"script"`
}

type CreateEndpointOptions struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Private   bool   `json:"private"`
	Method    string `json:"method"`
	// if exec id is emty a script id must be provide
	ScriptID string `json:"script_id"`
}

type CreateEndpointResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
	ID     string `json:"id"`
}

type GetEndpointResponse struct {
	Status string   `json:"status"`
	Data   Endpoint `json:"data"`
}

type GetEndpointListResponse struct {
	Status string     `json:"status"`
	Data   []Endpoint `json:"data"`
}

type DeleteEndpointResponse struct {
	Status string `json:"status"`
}

type ContextLookup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Namespace struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Env         []string `json:"env"`
	Context     string   `json:"context"`
}

type CreateNamespaceOptions struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Env         []string `json:"env"`
}

type CreateNamespaceResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
	ID     string `json:"id"`
}

type GetNamespaceResponse struct {
	Status string    `json:"status"`
	Error  string    `json:"error"`
	Data   Namespace `json:"data"`
}

type GetNamespaceListResponse struct {
	Status string      `json:"status"`
	Data   []Namespace `json:"data"`
}

type DeleteNamespaceResponse struct {
	Status string `json:"status"`
}
