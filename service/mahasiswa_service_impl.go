package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"web-golang/utils"
	"web-golang/repository"
	"web-golang/model/web"
	"web-golang/model/domain"
	"web-golang/exception"
)

type MahasiswaServiceImpl struct {
	MahasiswaRepository repository.MahasiswaRepository
	DB *sql.DB
	Validate *validator.Validate
}

func NewMahasiswaService(mahasiswaRepository repository.MahasiswaRepository, DB *sql.DB, validate *validator.Validate) MahasiswaService {
	return &MahasiswaServiceImpl {
		MahasiswaRepository: mahasiswaRepository,
		DB: DB,
		Validate: validate,
	}
}

func (service *MahasiswaServiceImpl) FindAll(ctx context.Context) []web.MahasiswaResponse {
	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	paramahasiswa := service.MahasiswaRepository.FindAll(ctx, tx)
	return utils.ToMahasiswaResponses(paramahasiswa)
}

func (service *MahasiswaServiceImpl) FindById(ctx context.Context, mahasiswaId int) web.MahasiswaResponse {
	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	mahasiswa, err := service.MahasiswaRepository.FindById(ctx, tx, mahasiswaId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return utils.ToMahasiswaResponse(mahasiswa)
}

func (service *MahasiswaServiceImpl) Create(ctx context.Context, request web.MahasiswaCreateRequest) web.MahasiswaResponse {
	err := service.Validate.Struct(request)
	utils.PanicIfError(err)
	
	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	mahasiswa := domain.Mahasiswa{
		Id: request.Id,
		Nama: request.Nama,
		NIM: request.NIM,
		IPK: request.IPK,
	}

	mahasiswa = service.MahasiswaRepository.Save(ctx, tx, mahasiswa)

	return utils.ToMahasiswaResponse(mahasiswa)
}

func (service *MahasiswaServiceImpl) Delete(ctx context.Context, mahasiswaId int) {
	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	mahasiswa, err := service.MahasiswaRepository.FindById(ctx, tx, mahasiswaId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.MahasiswaRepository.Delete(ctx, tx, mahasiswa)
}