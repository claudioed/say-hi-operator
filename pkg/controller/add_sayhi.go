package controller

import (
	"github.com/claudioed/say-hi-operator/pkg/controller/sayhi"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, sayhi.Add)
}
