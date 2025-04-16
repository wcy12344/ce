package main

import (
	"context"
	"fmt"
	stlog "log"
	"test/log"
	"test/registry"
	"test/service"
)

func main() {
	log.Run("distributed.log")
	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)
	r := registry.Registration {
		ServiceName: "logService",
		ServiceURL: serviceAddress,
	}
	ctx,err := service.Start(
		context.Background(),
		host,
		port,
		r,
		log.RegisterHandler,
	)
	if err != nil {
		stlog.Fatalln(err)
	}
	<- ctx.Done()
	fmt.Println("Shutting down log service.")
	
}
