package service

import (
	"context"
	"web-golang/model/web"
)

type MahasiswaService interface {
	FindAll(ctx context.Context) []web.MahasiswaResponse
	FindById(ctx context.Context, mahasiswaId int) web.MahasiswaResponse
	Create(ctx context.Context, request web.MahasiswaCreateRequest) web.MahasiswaResponse
	Delete(ctx context.Context, mahasiswaId int)
	Update(ctx context.Context, request web.MahasiswaUpdateRequest) web.MahasiswaResponse
}