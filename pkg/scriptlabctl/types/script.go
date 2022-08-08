package types

type CreateScriptOptions struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	FileName    string `json:"file_name"`
	FileContent string `json:"file_content"`
}

// requests types

type Script struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	FileName string `json:"file_name"`
}

type CreateScriptRequest struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	FileName    string `json:"file_name"`
	FileContent string `json:"file_content"`
}

type CreateScriptResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
	ID     string `json:"id"`
}

type UpdateScriptFileRequest struct {
	ID          string `json:"id"`
	FileContent string `json:"file_content"`
}

type UpdateScriptFileResponse struct {
	Status string `json:"status"`
}

type GetScriptResponse struct {
	Status string `json:"status"`
	Data   Script `json:"data"`
}

type GetScriptListResponse struct {
	Status string   `json:"status"`
	Data   []Script `json:"data"`
}
