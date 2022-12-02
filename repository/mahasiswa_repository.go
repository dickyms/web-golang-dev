package repository

import (
	"context"
	"database/sql"
	"web-golang/model/domain"
)

type MahasiswaRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Mahasiswa
}

