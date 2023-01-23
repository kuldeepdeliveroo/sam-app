package main

import "log"

type config struct {
	count int
	name  string
}

func main() {
	//lambda.Start(handler.Handle)

	cfg := config{
		count: 49,
		name:  "Kuldeep",
	}

	log.Printf("config: %v\n", cfg)
	update(&cfg)
	log.Printf("config: %v\n", *&cfg)

}

func update(cfg *config) {
	cfg.name = "Manisha"
}
