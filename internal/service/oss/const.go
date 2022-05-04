package oss

type ConfigStruct struct {
	Expiration string     `json:"expiration"`
	Conditions [][]string `json:"conditions"`
}

type PolicyToken struct {
	AccessKeyId    string `json:"accessid"`
	Host           string `json:"host"`
	Expire         int64  `json:"expire"`
	Signature      string `json:"signature"`
	Policy         string `json:"policy"`
	Directory      string `json:"dir"`
	Callback       string `json:"callback"`
	FileNamePrefix string `json:"fileNamePrefix"`
}

type CallbackParam struct {
	CallbackUrl      string `json:"callbackUrl"`
	CallbackBody     string `json:"callbackBody"`
	CallbackBodyType string `json:"callbackBodyType"`
}
