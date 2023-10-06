package dto

type ImageData struct {
	ID     int64  `json:"id"`
	UserId int64  `json:"user_id"`
	Image  string `json:"images"`
}
