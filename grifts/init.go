package grifts

import (
	"github.com/denisqsound/it_goroskop_generator/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
