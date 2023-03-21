package main

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	handler "simple_server/httpserver/server/handler_func"
)

func main() {
	http.HandleFunc("/info", handler.InfoHandler)

	err := http.ListenAndServe(":3000", nil)

	if errors.Is(err, http.ErrServerClosed) {
		log.Debug().Msg("server closed\n")
	} else if err != nil {
		log.Debug().Msg("error starting server\n")
		os.Exit(1)
	} else {
		fmt.Println("starting server at port number 3000")
	}
}
