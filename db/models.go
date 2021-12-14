package db

type Quote struct {
	//ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Author string `json:"author,omitempty" bson:"author,omitempty"`
	Quote  string `json:"quote,omitempty" bson:"quote,omitempty"`
}
