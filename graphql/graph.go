package main

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/shivanshsinghraghuvanshi/toll-collector/tolltax"
)

type Server struct {
	tolltaxClient *tolltax.Client
}

func NewGraphQLServer(tolltaxurl string) (*Server, error) {
	// tolltax Service
	t, err := tolltax.NewClient(tolltaxurl)
	if err != nil {
		return nil, err
	}

	return &Server{
		t,
	}, nil
}

func (s *Server) Mutation() MutationResolver {
	return &mutationResolver{
		server: s,
	}
}

func (s *Server) Query() QueryResolver {
	return &queryResolver{
		server: s,
	}
}

func (s *Server) Account() AccountResolver {
	return &accountResolver{
		server: s,
	}
}

func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: s,
	})
}