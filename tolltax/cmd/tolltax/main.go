package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/ssinghraghuvanshi/toll-collector/tolltax"
	"github.com/tinrab/retry"
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

	var r Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()

	log.Println("Listening on port 8080...")
	s := NewService(r)
	log.Fatal(ListenGRPC(s, 8080))
}
