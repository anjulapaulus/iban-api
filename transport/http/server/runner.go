package server

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/anjulapaulus/iban-api/app/config"
	"github.com/anjulapaulus/iban-api/transport/http/router"
)

// Run runs the http server.
func Run(cfg config.AppConfig) *http.Server {

	// initialize the router
	r := router.Init()

	srv := &http.Server{
		Addr: "0.0.0.0:" + strconv.Itoa(cfg.Port),

		// good practice to set timeouts to avoid Slowloris attacks
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,

		// pass the instance of gorilla/mux in
		Handler: r,
	}

	// run the server in a separate go routine so that it doesnt block
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			fmt.Println(err)
			panic("HTTP server shutting down unexpectedly...")
		}
	}()

	fmt.Printf("HTTP server listening on %v ...\n", srv.Addr)

	return srv
}

// Stop stops the server.
func Stop(ctx context.Context, srv *http.Server) {

	fmt.Println("HTTP server shutting down...")

	srv.Shutdown(ctx)
}
