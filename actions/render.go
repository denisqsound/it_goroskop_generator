package actions

import (
	"github.com/denisqsound/it_goroskop_generator/public"
	"github.com/denisqsound/it_goroskop_generator/templates"

	"github.com/gobuffalo/buffalo/render"
)

var r *render.Engine

func init() {
	r = render.New(render.Options{
		HTMLLayout:  "application.plush.html",
		TemplatesFS: templates.FS(),
		AssetsFS:    public.FS(),
		Helpers:     render.Helpers{},
	})
}
