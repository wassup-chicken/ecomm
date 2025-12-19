package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/wassup-chicken/jobs/internal/models"
)

//to query for specific usecases

func (s *store) GetUser(ctx context.Context, id int) (*models.User, error) {
	row := s.db.QueryRowContext(ctx, `select user_id, email from users where user_id = $1`, id)

	var user models.User

	err := row.Scan(&user.ID, &user.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("no results found")
		}
		return nil, err
	}

	return &user, nil
}

func (s *store) GetJob(ctx context.Context, id int) (*models.Job, error) {
	row := s.db.QueryRowContext(ctx, `select id, title, description from jobs where id = $1`, id)

	var job models.Job

	err := row.Scan(&job.ID, &job.Title, &job.Description)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no result with id %d", id)
		}
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
