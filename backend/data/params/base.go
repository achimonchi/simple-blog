package params

type Response struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
	Token   string `json:"token,omitempty"`
}
