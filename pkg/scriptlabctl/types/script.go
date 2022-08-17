package types

import "time"

type CreateScriptOptions struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	FileName    string `json:"file_name"`
	FileContent string `json:"file_content"`
}

// requests types

type Script struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	FileName    string    `json:"file_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateScriptRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
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
	Status    string    `json:"status"`
	Error     string    `json:"error"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetScriptResponse struct {
	Status string `json:"status"`
	Data   Script `json:"data"`
}

type GetScriptListResponse struct {
	Status string   `json:"status"`
	Error  string   `json:"error"`
	Data   []Script `json:"data"`
}
