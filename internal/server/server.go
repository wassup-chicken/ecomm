package server

import (
	"github.com/wassup-chicken/jobs/internal/clients"
)

type Server struct {
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

	if err != nil {
		return nil, err
	}

	return &Server{
		Firebase: app,
		LLM:      openai,
	}, nil
}
