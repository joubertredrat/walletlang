package dao

type User struct {
	ID     string `bson:"_id,omitempty"`
	AppID  string `bson:"app_id"`
	Amount uint   `bson:"amount"`
}
