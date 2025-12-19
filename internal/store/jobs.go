package store

import (
	"context"

	"github.com/wassup-chicken/jobs/internal/models"
)

//to query for specific usecases

func (s *store) GetJob(ctx context.Context, id string) (*models.Job, error) {
	row := s.db.QueryRowContext(ctx, `select id, title, description from jobs where id = $1`, id)

	var job models.Job

	err := row.Scan(&job.ID, &job.Title, &job.Description)

	if err != nil {
		return nil, err
	}

	return &job, nil
}

func (s *store) GetJobs(ctx context.Context) (*[]models.Job, error) {
	q := `select id, title, description from jobs`
	rows, err := s.db.QueryContext(ctx, q)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []models.Job
	for rows.Next() {
		var job models.Job
		err := rows.Scan(&job.ID, &job.Title, &job.Description)
		if err != nil {
			return nil, err
		}

		jobs = append(jobs, job)
	}

	return &jobs, nil
}
