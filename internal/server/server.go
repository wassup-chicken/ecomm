package server

import (
	"github.com/wassup-chicken/jobs/internal/clients"
	store "github.com/wassup-chicken/jobs/internal/store"
)

type JobServer struct {
	JobStore store.JobStorer
	Firebase Auth
	LLM      clients.LLM
}

func New() (*JobServer, error) {
	//initialize firebase app
	app, err := NewAUth()

	if err != nil {
		return nil, err
	}

	//initialize clients
	openai := clients.NewLLM()

	//initialize database
	store, err := store.NewStore()

	if err != nil {
		return nil, err
	}

	return &JobServer{
		JobStore: store,
		Firebase: app,
		LLM:      openai,
	}, nil
}
