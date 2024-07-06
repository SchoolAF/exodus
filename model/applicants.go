package model

type MaintainerShipForm struct {
	ID               string `json:"id"`
	Name             string `json:"name" valid:"required"`
	Username         string `json:"username" valid:"required"`
	Email            string `json:"email" valid:"required"`
	Telegram         string `json:"telegram" valid:"required"`
	Github           string `json:"github" valid:"required"`
	Gitlab           string `json:"gitlab" valid:"required"`
	Brand            string `json:"brand" valid:"required"`
	Marketname       string `json:"marketname" valid:"required"`
	Codename         string `json:"codename" valid:"required"`
	DeviceTree       string `json:"devicetreeurl" valid:"required"`
	CommonTree       string `json:"commontreeurl"`
	Kernel           string `json:"kernelurl" valid:"required"`
	VendorTree       string `json:"vendortreeurl" valid:"required"`
	VendorCOmmonTree string `json:"devicecommonurl"`
	Status           string `json:"status"`
}
