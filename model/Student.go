package model

type Student struct {
	Id             int    `json:"Id,omitempty"`
	Name           string `json:"Name,omitempty"`
	Birthday       string `json:"Birthday"`
	Grade          int    `json:"Grade"`
	Gender         string `json:"Gender"`
	PhoneNum       string `json:"PhoneNum"`
	Email          string `json:"Email"`
	School         string `json:"School"`
	ParentName     string `json:"ParentName"`
	ParentPhoneNum string `json:"ParentPhoneNum"`
}
