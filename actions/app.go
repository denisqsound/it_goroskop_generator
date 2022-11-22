package actions

import (
	"github.com/denisqsound/it_goroskop_generator/models"
	"github.com/gobuffalo/buffalo-pop/v3/pop/popmw"
	"net/http"

	"github.com/denisqsound/it_goroskop_generator/locales"
	"github.com/denisqsound/it_goroskop_generator/public"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/unrolled/secure"

	csrf "github.com/gobuffalo/mw-csrf"
	forcessl "github.com/gobuffalo/mw-forcessl"
	i18n "github.com/gobuffalo/mw-i18n/v2"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
)

var ENV = envy.Get("GO_ENV", "development")

var (
	app *buffalo.App
	T   *i18n.Translator
)

func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "it_goroskop_generator_session",
		})

		app.Use(paramlogger.ParameterLogger)
		app.Use(csrf.New)
		app.Use(popmw.Transaction(models.DB))
		app.Use(translations())

		app.GET("/", HomeHandler)
		app.GET("/money", MoneyHandler)

		app.GET("/iamgay", PredictionHandler)

		app.ServeFiles("/", http.FS(public.FS())) // serve files from the public directory
	}

	return app
}

func translations() buffalo.MiddlewareFunc {
	var err error
	if T, err = i18n.New(locales.FS(), "en-US"); err != nil {
		app.Stop(err)
	}
	return T.Middleware()
}

func forceSSL() buffalo.MiddlewareFunc {
	return forcessl.Middleware(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}
