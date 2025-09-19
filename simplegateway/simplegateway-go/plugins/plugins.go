package plugins

import (
	"github.com/jzero-io/jzero/core/restc"
)

type Plugins any

func NewPlugins(c restc.Client) Plugins {
	return &plugins{}
}

type plugins struct {
}
