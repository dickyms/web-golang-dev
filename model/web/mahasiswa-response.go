package web

type MahasiswaResponse struct {
	Id int `json:"id"`
	Nama string `json:"name"`
	NIM int `json:"nim"`
	IPK int `json:"ipk"`
}