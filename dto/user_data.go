package dto

type UserData struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Username string `json:"username"`
	Password string `json:"password"`
}
