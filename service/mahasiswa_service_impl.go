package service

import (
	"context"
	"database/sql"
	"web-golang/utils"
	"web-golang/repository"
	"web-golang/model/web"
	"web-golang/exception"
)

type MahasiswaServiceImpl struct {
	MahasiswaRepository repository.MahasiswaRepository
	DB *sql.DB
}

func NewMahasiswaService(mahasiswaRepository repository.MahasiswaRepository, DB *sql.DB) MahasiswaService {
	return &MahasiswaServiceImpl {
		MahasiswaRepository: mahasiswaRepository,
		DB: DB,
	}
}

func (service *MahasiswaServiceImpl) FindAll(ctx context.Context) []web.MahasiswaResponse {
	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	paramahasiswa := service.MahasiswaRepository.FindAll(ctx, tx)
	return utils.ToMahasiswaResponses(paramahasiswa)
}

func (service *MahasiswaServiceImpl) FindAll(ctx context.Context, mahasiswaId int) web.MahasiswaResponse {
	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	mahasiswa, err := service.MahasiswaRepository.FindById(ctx, tx, mahasiswaId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
}