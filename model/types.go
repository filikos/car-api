package model

type Car struct {
	ID      string `json:"id"`
	Model   string `json:"model"`
	Make    string `json:"make"`
	Variant string `json:"variant"`
	// TODO: Add car properties
}

type Cars struct {
	Cars []Car `json:"cars"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
