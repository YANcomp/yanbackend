package main

import (
	"context"
	"flag"
	"github.com/YANcomp/yanbackend/stories_service/internal/app"
	"log"
)

func main() {
	ctx := context.Background()

	flag.Parse()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
