package model

type Version struct {
	Codename       string `json:"codename" valid:"required"`
	Branch         string `json:"branch" valid:"required"`
	AndroidVersion string `json:"android_version" valid:"required"`
}
