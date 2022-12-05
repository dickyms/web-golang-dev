package web

type MahasiswaUpdateRequest struct {
	Id int `validate:"required"`
	Nama string `validate:"required,min=1,max=100" json:"Nama"`
	NIM int `validate:"required" json:"NIM"`
	IPK int `validate:"required" json:"IPK"`
}