package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/anjulapaulus/iban-api/app/config"
	httpServer "github.com/anjulapaulus/iban-api/transport/http/server"
)

func main() {
	// parse all configurations
	cfg := config.Parse("./configs")

	fmt.Println("Service starting...")

	// start the server to handle http requests
	hsrv := httpServer.Run(cfg.AppConfig)

	fmt.Println("Ready")

	// enable graceful shutdown
	c := make(chan os.Signal, 1)

	// accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught
	signal.Notify(c, os.Interrupt)

	// block until a registered signal is received
	<-c

	// Shutdown in the reverse order of initialization.
	fmt.Println("\nService stopping...")

	// create a deadline to wait for
	var wait time.Duration

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// gracefully stop the http server
	httpServer.Stop(ctx, hsrv)

	fmt.Println("Done")

	os.Exit(0)

}
