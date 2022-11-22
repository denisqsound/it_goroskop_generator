package actions

import (
	"github.com/gobuffalo/buffalo"
	"net/http"
)

func PredictionHandler(ctx buffalo.Context) error {
	return ctx.Render(http.StatusOK, r.JSON(map[string]string{"message": "YOU GAY!"}))
}
