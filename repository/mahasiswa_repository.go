package repository

import (
	"context"
	"database/sql"
	"web-golang/model/domain"
)

type MahasiswaRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Mahasiswa
	FindById(ctx context.Context, tx *sql.Tx, mahasiswaId int) (domain.Mahasiswa, error)
	Save(ctx context.Context, tx *sql.Tx, mahasiswa domain.Mahasiswa) domain.Mahasiswa
	Delete(ctx context.Context, tx *sql.Tx, mahasiswa domain.Mahasiswa)
	Update(ctx context.Context, tx *sql.Tx, mahasiswa domain.Mahasiswa) domain.Mahasiswa
}

