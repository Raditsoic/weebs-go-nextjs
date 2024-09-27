package model

type UserDB struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Salt      string `json:"salt"`
	Hash      string `json:"hash"`
	CreatedAt string `json:"createdAt"`
}
