package kernel

import (
	"todo/pkg"
	"todo/pkg/logger"
)

type stopper struct {
	*pkg.Server
}

func (s *stopper) Stop() {
	if s.DB != nil {
		_ = s.DB.Close()
	}
}

func Bootstrap(serviceName string, server *pkg.Server) *stopper {
	log := logger.New(serviceName).WithServiceInfo("Bootstrap")
	log.Println("starting...")

	if server.Http != nil {
		address := server.HttpAddress
		log.Printf("http address: %s", address)
		go func() {
			MustServeHttp(server.Http, address)
		}()
	}

	log.Info("server is ready")

	s := &stopper{
		Server: server,
	}

	return s
}
