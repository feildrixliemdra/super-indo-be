package payload

type Response struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Error   any               `json:"error,omitempty"`
	Errors  []ErrorValidation `json:"errors,omitempty"`
	Data    any               `json:"data,omitempty"`
}

type ListResponse struct {
	Success     bool        `json:"success"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
	TotalData   int64       `json:"total_data"`
	TotalPage   int64       `json:"total_page"`
	CurrentPage int64       `json:"current_page"`
}

type ErrorValidation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
