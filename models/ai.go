package models

type RoomRequest struct {
	Person   string `json:"person"`
	Facility string `json:"facility"`
}

type AIResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}
