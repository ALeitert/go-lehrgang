package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/risingwavelabs/eris"

	"weather-service/internal/config"
)

type Server struct {
	server *http.Server
}

func (Server) Name() string { return "API Server" }

func (svr *Server) Init(ctx context.Context) error {
	router := http.NewServeMux()
	router.HandleFunc("GET /", get)

	svr.server = &http.Server{
		Addr:    ":" + strconv.Itoa(int(config.C.APIPort)),
		Handler: router,
	}

	return nil
}

func (svr *Server) Run(ctx context.Context) error {
	fmt.Printf("%s is listening on port %s\n", svr.Name(), svr.server.Addr)

	svr.server.BaseContext = func(_ net.Listener) context.Context {
		return ctx
	}

	err := svr.server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return eris.Wrapf(err, "%s stopped", svr.Name())
	}

	return nil
}

func (svr *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := svr.server.Shutdown(ctx)
	if err != nil {
		return eris.Wrapf(err, "failed to shut down %s", svr.Name())
	}

	return nil
}
