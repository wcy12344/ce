package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"test/registry"
)

func main() {
	http.Handle("/services",&registry.RegistryService{})

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var srv http.Server
	srv.Addr = registry.ServerPort
	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func(){
		fmt.Print("Registry Service started. Press and key to stop.\n")
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()

	<- ctx.Done()
	fmt.Print("Shutting down registry service.")
}
