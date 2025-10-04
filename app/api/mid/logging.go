package mid

import (
	"context"
	"net/http"

	"github.com/nasissa97/service/foundation/logger"
	"github.com/nasissa97/service/foundation/web"
)

func Logger(log *logger.Logger) web.MidHandler {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			log.Info(ctx, "request started", "method", r.Method, "path", r.URL.Path, "remote", r.RemoteAddr)
			err := handler(ctx, w, r)
			log.Info(ctx, "request completed", "method", r.Method, "path", r.URL.Path, "remote", r.RemoteAddr)
			return err
		}
		return h
	}

	return m
}
