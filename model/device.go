package model

type Brand struct {
	Brand      string `json:"brand" valid:"required"`
	BrandLower string `json:"brand_lower" valid:"required"`
}

type Device struct {
	Marketname string `json:"marketname" valid:"required"`
	Codename   string `json:"codename" valid:"required"`
	BrandLower string `json:"brand_lower" valid:"required"`
	Deprecated bool   `json:"deprecated"`
	Maintainer string `json:"maintainer" valid:"required"`
}
