package dao

import "time"

type User struct {
	ID        string     `bson:"_id,omitempty"`
	AppID     string     `bson:"app_id"`
	Amount    uint       `bson:"amount"`
	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt *time.Time `bson:"updated_at"`
}
