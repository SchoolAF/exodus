package model

type Maintainer struct {
	Username string `json:"username" valid:"required"`
	Name     string `json:"name" valid:"required"`
	GitHub   string `json:"github" valid:"required"`
	GitLab   string `json:"gitlab" valid:"required"`
	Telegram string `json:"telegram" valid:"required"`
	Email    string `json:"email" valid:"required"`
}
