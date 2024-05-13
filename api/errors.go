package api

import (
	"fmt"
	"net/http"
)

type APIError struct {
	Status  int `json:"status"`
	Details any `json:"details"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("Status code: %d, Details: %s", e.Status, e.Details)
}

func InvalidData(errors map[string]any) *APIError {
	return &APIError{
		Status:  http.StatusBadRequest,
		Details: errors,
	}
}
