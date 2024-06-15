package dto

type CommandRequest struct {
	CMD bool `json:"CMD"`
	Num int  `json:"num"`
}

type CommandResponse struct {
	Message string `json:"message"`
}
