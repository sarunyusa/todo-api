package kernel

import (
	"net/http"
	"todo/pkg/logger"
)

func MustServeHttp(handler http.Handler, address string) {
	log := logger.New("starter").WithServiceInfo("MustServeHttp")
	srv := &http.Server{
		Addr:    address,
		Handler: handler,
	}

	log.Infof("start serving HTTP on %s", address)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.WithError(err).Fatalf("http server error")
	}
}
