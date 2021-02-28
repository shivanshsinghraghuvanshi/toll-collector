package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/shivanshsinghraghuvanshi/toll-collector/tolltax"
	"github.com/tinrab/retry"
)

type Config struct {
	DatabaseURL       string `envconfig:"DATABASE_URL"`
	PaymentServiceURL string `envconfig:"PAYMENT_SERVICE_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var r tolltax.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = tolltax.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()

	log.Println("Listening on port 8080...")
	s := tolltax.NewService(r)
	log.Fatal(tolltax.ListenGRPC(s, cfg.PaymentServiceURL, 8080))
}
