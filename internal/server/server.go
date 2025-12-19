package server

import store "github.com/wassup-chicken/jobs/internal/store"

type JobServer struct {
	jobStore store.JobStorer
}

func New() (*JobServer, error) {
	//initialize database
	store, err := store.New()

	if err != nil {
		return nil, err
	}

	return &JobServer{
		jobStore: store,
	}, nil
}
