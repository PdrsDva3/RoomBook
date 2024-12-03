package models

type Tag struct {
	id   int    `json:"id"`
	Type string `json:"type"`
	Text string `json:"text"`
}
