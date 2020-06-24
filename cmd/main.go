package main

import (
	"todo/controller/todo"
	"todo/internal/kernel"
	"todo/pkg/logger"
)

func main() {
	serviceName := "todo"
	log := logger.New(serviceName)

	log.Info("starting service")

	config := kernel.LoadConfig()
	config.ConnectDB()
	config.MustPrepareDatabase()

	s := make([]kernel.Stopper, 0)
	s = append(s, kernel.Bootstrap("todo", todo.New(config)))
	kernel.GracefullyStop(serviceName, s...)
}
