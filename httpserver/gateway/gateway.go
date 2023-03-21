package gateway

import (
	"errors"
	"github.com/rs/zerolog/log"
	"net/http"
	handler "simple_server/httpserver/server/handler_func"
)

func main() {
	http.HandleFunc("/info", handler.InfoHandler)

	err := http.ListenAndServe(":3000", nil)

	if errors.Is(err, http.ErrServerClosed) {
		log.Debug().Msg("server closed\n")
	} else {
		log.Debug().Msg("error starting server\n")
	}
}
