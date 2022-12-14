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
	SQL := "select Id, Nama, NIM, IPK from mahasiswa where Id = $1"
	rows, err := tx.QueryContext(ctx, SQL, mahasiswaId)
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

func (repository *MahasiswaRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, mahasiswa domain.Mahasiswa) domain.Mahasiswa {
	SQL := "INSERT INTO mahasiswa VALUES($1, $2, $3, $4)"
	_, err := tx.ExecContext(ctx, SQL, mahasiswa.Id, mahasiswa.Nama, mahasiswa.NIM, mahasiswa.IPK)
	utils.PanicIfError(err)

	return mahasiswa
}

func (repository *MahasiswaRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, mahasiswa domain.Mahasiswa) {
	SQL := "DELETE FROM mahasiswa WHERE Id = $1"
	_, err := tx.ExecContext(ctx, SQL, mahasiswa.Id)
	utils.PanicIfError(err)
}

func (repository *MahasiswaRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, mahasiswa domain.Mahasiswa) domain.Mahasiswa {
	SQL := "UPDATE mahasiswa set Nama = $2, NIM = $3, IPK = $4 WHERE Id = $1"
	_, err := tx.ExecContext(ctx, SQL, mahasiswa.Id, mahasiswa.Nama, mahasiswa.NIM, mahasiswa.IPK)
	utils.PanicIfError(err)

	return mahasiswa
}