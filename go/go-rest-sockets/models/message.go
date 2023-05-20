package models

type WebsocketMessage struct {
	Type string `json:"type"`
	Payload any `json:"payload"`
}