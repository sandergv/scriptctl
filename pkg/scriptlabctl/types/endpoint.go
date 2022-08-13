package types

type NamespaceLookup struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Endpoint struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
	Private   bool   `json:"private"`
	Method    string `json:"method"`
	ExecID    string `json:"exec_id"`
}

type CreateEndpointOptions struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Private   bool   `json:"private"`
	Method    string `json:"method"`

	// if ExecID is empty a new exec it will be created
	ExecID string `json:"exec_id"`

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

type Namespace struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Endpoints []Endpoint `json:"endpoints"`
}

type CreateNamespaceOptions struct {
	Name string   `json:"name"`
	Env  []string `json:"env"`
}

type CreateNamespaceResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
	ID     string `json:"id"`
}

type GetNamespaceResponse struct {
	Status string    `json:"status"`
	Data   Namespace `json:"data"`
}

type GetNamespaceListResponse struct {
	Status string      `json:"status"`
	Data   []Namespace `json:"data"`
}

type DeleteNamespaceResponse struct {
	Status string `json:"status"`
}
