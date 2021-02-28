package main

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
)

type AppConfig struct {
	TollTaxServiceURL string `envconfig:"TOLLTAX_SERVICE_URL"`
	PaymentServiceURL string `envconfig:"PAYMENT_SERVICE_URL"`
}

func main() {
	var cfg AppConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v context is loaded from env\n", cfg.TollTaxServiceURL)
	s, err := NewGraphQLServer(cfg.TollTaxServiceURL, cfg.PaymentServiceURL)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/graphql", handler.GraphQL(s.ToExecutableSchema()))
	http.Handle("/playground", handler.Playground("TollCollector", "/graphql"))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
