package main

import (
	"apigateway/internal/container"
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	c, err := container.NewContainer(ctx)
	if err != nil {
		panic(err)
	}
	log.Println(c.GetConfig().JWT.SecretKey)
}
