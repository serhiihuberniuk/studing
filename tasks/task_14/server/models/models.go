package models

import (
	"encoding/json"
	"fmt"
)

type RequestMessage struct {
	Text string `json:"text"`
}

func (m *RequestMessage) ToBytes() ([]byte, error) {
	jsonMessage, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("error while marshling response message: %w", err)
	}

	return jsonMessage, nil
}

func (m *RequestMessage) FromBytes(data []byte) error {
	if err := json.Unmarshal(data, m); err != nil {
		return fmt.Errorf("error while unmarshalling response message: %w", err)
	}

	return nil
}

type ResponseMessage struct {
	Success bool   `json:"success"`
	Text    string `json:"text"`
}

func (m *ResponseMessage) ToBytes() ([]byte, error) {
	jsonMessage, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("error while marshling response message: %w", err)
	}

	return jsonMessage, nil
}

func (m *ResponseMessage) FromBytes(data []byte) error {
	if err := json.Unmarshal(data, m); err != nil {
		return fmt.Errorf("error while unmarshalling response message: %w", err)
	}

	return nil
}
