package web

type MahasiswaCreateRequest struct {
	Id int
	Nama string `validate:"required,min=1,max=100" json:"Nama"`
	NIM int `validate:"required" json:"NIM"`
	IPK int `validate:"required" json:"IPK"`
}