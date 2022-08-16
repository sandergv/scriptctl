package types

import "time"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Status      string    `json:"status"`
	Error       string    `json:"error"`
	WorkspaceID string    `json:"workspace_id"`
	Token       string    `json:"token"`
	ExpiresAt   time.Time `json:"expires_at"`
}

type AuthDetails struct {
	WorkspaceID string    `json:"workspace_id"`
	Token       string    `json:"token"`
	ExpiresAt   time.Time `json:"expires_at"`
}
