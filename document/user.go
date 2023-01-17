package document

import "time"

type UserDocument struct {
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Married   bool      `json:"married"`
	CreatedAt time.Time `json:"createdAt"`
	About     string    `json:"about"`
}
