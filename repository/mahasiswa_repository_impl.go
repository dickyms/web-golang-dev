package repository

import (
	"errors"
	"context"
	"database/sql"
	"web-golang/utils"
	"web-golang/model/domain"
)

type MahasiswaRepositoryImpl struct {

}

func NewMahasiswaRepository() MahasiswaRepository {
	return &MahasiswaRepositoryImpl{}
}

func (repository *MahasiswaRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Mahasiswa {
	SQL := "select * from mahasiswa"
	rows, err := tx.QueryContext(ctx, SQL)
	utils.PanicIfError(err)
	defer rows.Close()

	var paramahasiswa []domain.Mahasiswa
	for rows.Next() {
		mahasiswa := domain.Mahasiswa{}
		err := rows.Scan(&mahasiswa.Id, &mahasiswa.Nama, &mahasiswa.NIM, &mahasiswa.IPK)
		utils.PanicIfError(err)
		paramahasiswa = append(paramahasiswa, mahasiswa)
	}
	return paramahasiswa
}

func (repository *MahasiswaRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, mahasiswaId int) (domain.Mahasiswa, error) {
	SQL := "SELECT * FROM mahasiswa WHERE Id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	utils.PanicIfError(err)
	defer rows.Close()

	mahasiswa := domain.Mahasiswa{}
	if rows.Next() {
		err := rows.Scan(&mahasiswa.Id, &mahasiswa.Nama, &mahasiswa.NIM, &mahasiswa.IPK)
		utils.PanicIfError(err)
		return mahasiswa, nil
	} else {
		return mahasiswa, errors.New("Mahasiswa is not found")
	}
}
