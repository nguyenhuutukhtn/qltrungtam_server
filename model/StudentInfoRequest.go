package model

type StudentInfoRequest struct {
	id            int    `json:"id,omitempty"`
	HoTen         string `json:"HoTen,omitempty"`
	NgaySinh      string `json:"NgaySinh"`
	Lop           int    `json:"Lop"`
	GioiTinh      string `json:"GioiTinh"`
	SDT           string `json:"SDT"`
	Email         string `json:"Email"`
	Truong        string `json:"Truong"`
	SDTPhuHuynh   string `json:"SDTPhuHunh"`
	HoTenPhuHuynh string `json:"HoTenPhuHuynh"`
}
