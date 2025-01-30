package types

type (
	ValidationError struct {
		Field string      `json:"field"` // The field name that failed validation
		Tag   string      `json:"tag"`   // The validation tag that failed
		Value interface{} `json:"value"` // The invalid value
		Error bool        `json:"error"` // Indicates if an error occurred
	}

	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)
