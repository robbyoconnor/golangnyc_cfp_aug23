package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/robbyoconnor/golangnyc_cfp_aug23/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
