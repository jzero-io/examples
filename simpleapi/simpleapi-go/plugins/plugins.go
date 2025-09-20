package plugins

import (
	"github.com/jzero-io/jzero/core/restc"
)

type Plugins any

type plugins struct {
}

func NewPlugins(c restc.Client) Plugins {
	return &plugins{}
}
