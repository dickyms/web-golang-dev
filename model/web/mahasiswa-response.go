package web

type MahasiswaResponse struct {
	Id int `json:"id"`
	Nama string `json:"nama"`
	NIM int `json:"nim"`
	IPK int `json:"ipk"`
}