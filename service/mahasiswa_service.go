package service

import (
	"context"
	"web-golang/model/web"
)

type MahasiswaService interface {
	FindAll(ctx context.Context) []web.MahasiswaResponse
	FindById(ctx context.Context, mahasiswaId int) web.MahasiswaResponse
}