package dto

type ControlRequest struct {
	CMD bool `json:"CMD"`
}

type ControlResponse struct {
	Message string `json:"message"`
}
