// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"os"

	"github.com/nasissa97/service/api/services/api/mid"
	"github.com/nasissa97/service/api/services/sales/route/sys/checkapi"
	"github.com/nasissa97/service/foundation/logger"
	"github.com/nasissa97/service/foundation/web"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(log *logger.Logger, shutdown chan os.Signal) *web.App {
	mux := web.NewApp(shutdown, mid.Logger(log))

	checkapi.Routes(mux)

	return mux
}
