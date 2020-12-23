package users

import "time"

type User struct {
	ID       string    `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Status   bool      `json:"status"`
	Updated  time.Time `json:"updated"`
	Created  time.Time `json:"created"`
}
