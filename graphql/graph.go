package main

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/shivanshsinghraghuvanshi/toll-collector/graphql/graph/generated"
	"github.com/shivanshsinghraghuvanshi/toll-collector/payment"
	"github.com/shivanshsinghraghuvanshi/toll-collector/tolltax"
	"log"
)

type Server struct {
	tolltaxClient *tolltax.Client
	paymentClient *payment.Client
}

func NewGraphQLServer(tolltaxurl string, paymentServiceUrl string) (*Server, error) {
	log.Println("tolltax url and payment service url")
	// TOLL TAX SERVICE LOADING
	t, err := tolltax.NewClient(tolltaxurl)
	if err != nil {
		return nil, err
	}

	// Payment Service Wiring
	p, e := payment.NewClient(paymentServiceUrl)
	if e != nil {
		return nil, err
	}
	return &Server{
		t, p,
	}, nil
}

func (s *Server) Mutation() generated.MutationResolver {
	return mutationResolver{server: s}
}

func (s *Server) Query() generated.QueryResolver {
	return queryResolver{server: s}
}

func (s *Server) TollTax() *tolltaxResolver {
	return &tolltaxResolver{
		server: s,
	}
}

func (s *Server) Payment() *paymentResolver {
	return &paymentResolver{
		server: s,
	}
}

func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: s,
	})
}
