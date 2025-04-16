package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"test/registry"
)

func Start(ctx context.Context, host, port string, reg registry.Registration,
	registryHandlersFunc func()) (context.Context, error) {
	registryHandlersFunc()
	ctx = StartService(ctx, reg.ServiceName , host, port)
	err := registry.RegisterService(reg)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}

func StartService(ctx context.Context, serviceName registry.ServiceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server
	srv.Addr = ":" + port

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Print("%v started. Press and key to stop.\n", serviceName)
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()

	return ctx
}
