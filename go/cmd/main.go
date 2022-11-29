package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/darrensemusemu/protohackers/go/internal/smoketest"
	"github.com/darrensemusemu/protohackers/go/pkg/server"
)

type Config struct {
	port    int
	problem string
}

func main() {
	cfg := Config{}
	flag.IntVar(&cfg.port, "port", 8080, "Listening port")
	flag.StringVar(&cfg.problem, "problem", "-0", "Problem solution to run")
	flag.Parse()

	if err := run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}

func run(cfg Config) error {
	var svc server.Service

	switch cfg.problem {
	case "-0":
		svc = smoketest.New()
	default:
		return fmt.Errorf("'problem' flag'%s' not valid, see -help", cfg.problem)
	}

	s := server.New(svc)
	if err := s.Run(cfg.port); err != nil {
		return err
	}

	return nil
}
