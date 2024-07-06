package model

type Accounts struct {
	ID       string `json:"id" valid:"required"`
	Username string `json:"username" valid:"required"`
	Passkey  int    `json:"passkey"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username" valid:"required"`
	Passkey  int    `json:"passkey" valid:"required"`
}
