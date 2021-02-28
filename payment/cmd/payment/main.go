package main

import (
	"github.com/shivanshsinghraghuvanshi/toll-collector/payment"
	"github.com/tinrab/retry"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var r payment.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = payment.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()

	log.Println("Listening on port 8080...")
	s := payment.NewService(r)
	log.Fatal(payment.ListenGRPC(s, 8080))
}
