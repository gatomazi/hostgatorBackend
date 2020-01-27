package models

// ErrorModel -
type ErrorModel struct {
	Type        string `json:"tipo"`
	Status      int    `json:"status"`
	Message     string `json:"message"`
	ResultError string `json:"resultError"`
}

// MakeNewErrorResponse -
func MakeNewErrorResponse(typ string, status int, message, resultError string) ErrorModel {
	return ErrorModel{Type: typ, Status: status, Message: message, ResultError: resultError}
}
