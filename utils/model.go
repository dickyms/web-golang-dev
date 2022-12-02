package utils

import (
	"web-golang/model/domain"
	"web-golang/model/web"
)

func ToMahasiswaResponse(mahasiswa domain.Mahasiswa) web.MahasiswaResponse {
	return web.MahasiswaResponse{
		Id: mahasiswa.Id,
		Nama: mahasiswa.Nama,
		NIM: mahasiswa.NIM,
		IPK: mahasiswa.IPK,
	}
}

func ToMahasiswaResponses(paramahasiswa []domain.Mahasiswa) []web.MahasiswaResponse {
	var mahasiswaResponses []web.MahasiswaResponse
	for _, mahasiswa := range paramahasiswa {
		mahasiswaResponses = append(mahasiswaResponses, ToMahasiswaResponse(mahasiswa))
	}
	return mahasiswaResponses
}