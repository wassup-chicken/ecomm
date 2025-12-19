package store

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/wassup-chicken/jobs/internal/models"
)

// to initialize db connections

type store struct {
	db *sql.DB
}

type JobStorer interface {
	GetJob(ctx context.Context, id string) (*models.Job, error)
	GetJobs(ctx context.Context) (*[]models.Job, error)
}

// initialize the database conection and returns a repo instance
func New() (JobStorer, error) {
	db, err := sql.Open("pgx", "postgres://shong@localhost:5432/jobs")

	if err != nil {
		return nil, err
	}

	return &store{
		db: db,
	}, nil

}
