package models

type RequestMessage struct {
	Text string `json:"text"`
}

type ResponseMessage struct {
	Success bool   `json:"success"`
	Text    string `json:"text"`
}
