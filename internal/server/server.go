package server

import (
	"github.com/wassup-chicken/jobs/internal/clients"
	store "github.com/wassup-chicken/jobs/internal/store"
)

type Server struct {
	JobStore store.JobStorer
	Firebase Auth
	LLM      clients.LLM
}

func New() (*Server, error) {
	//initialize firebase app
	app, err := NewAuth()

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

	return &Server{
		JobStore: store,
		Firebase: app,
		LLM:      openai,
	}, nil
}
