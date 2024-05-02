package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nurzzaat/ZharasDiplom/internal/models"
)

type SyllabusRepository struct {
	db *pgxpool.Pool
}

func NewSyllabusRepository(db *pgxpool.Pool) models.SyllabusRepository {
	return &SyllabusRepository{db: db}
}

func (sr *SyllabusRepository) Create(c context.Context) error {
	return nil
}
