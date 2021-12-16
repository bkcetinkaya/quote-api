package db

type Quote struct {
	Author string `json:"author,omitempty" bson:"author,omitempty"`
	Quote  string `json:"quote,omitempty" bson:"quote,omitempty"`
}
