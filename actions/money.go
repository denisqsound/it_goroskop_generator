package actions

import (
	"github.com/gobuffalo/buffalo"
	"net/http"
)

func MoneyHandler(ctx buffalo.Context) error {
	return ctx.Render(http.StatusOK, r.HTML("home/money.plush.html"))

}
