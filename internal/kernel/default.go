package kernel

import (
	"os"
	"os/signal"
	"syscall"
	"todo/pkg/logger"
)

func LoadConfig() *ConfigFile {
	return loadDevConfigFile()
}

type Stopper interface {
	Stop()
}

func GracefullyStop(serviceName string, stoppers ...Stopper) {
	log := logger.New(serviceName).WithServiceInfo("gracefullyStop")
	gracefullyClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGTERM)
		signal.Notify(sigint, syscall.SIGINT)
		sig := <-sigint

		log.Printf("‍system signal received: %s", sig.String())
		// stop service
		log.Info("stopping services")
		for _, s := range stoppers {
			s.Stop()
		}
		close(gracefullyClosed)
	}()
	<-gracefullyClosed
	log.Info("services stopped")
}
