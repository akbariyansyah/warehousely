package goods

import "time"

type Goods struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Stock      string    `json:"stock"`
	Status     bool      `json:"status"`
	CategoryID string    `json:"category_id"`
	UserID     string    `json:"user_id"`
	ImagePath  string    `json:"image_path"`
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated"`
}
