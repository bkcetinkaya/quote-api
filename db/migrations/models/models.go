package models

type Quote struct {
	ID     int64  `json:"id"`
	Author string `json:"author"`
	Quote  string `json:"quote"`
}
