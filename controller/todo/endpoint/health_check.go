package endpoint

import (
	"context"
	"net/http"
)

func HealthCheck(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(200)
	_, err := w.Write([]byte("TODO API"))
	return err
}
