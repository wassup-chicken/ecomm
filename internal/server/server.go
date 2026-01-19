package server

import (
	"github.com/wassup-chicken/jobs/internal/clients"
)

type Server struct {
	Firebase Auth
	LLM      clients.LLM
}

func New() (*Server, error) {
	//initialize clients

	//initialize firebase app
	app, err := NewAuth()

	if err != nil {
		return nil, err
	}

	//initialize clients
	openai := clients.NewLLM()

	return &Server{
		Firebase: app,
		LLM:      openai,
	}, nil
}
