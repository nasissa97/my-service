package checkapi

import (
	"github.com/nasissa97/service/foundation/web"
)

func Routes(app *web.App) {
	app.HandleFunc("GET /liveness", liveness)
	app.HandleFunc("GET /readiness", readiness)
}
