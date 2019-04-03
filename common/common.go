package common

import (
	"github.com/JayneJacobs/gopherclub/common/datastore"

	"go.isomorphicgo.org/go/isokit"
)

type Env struct {
	DB          datastore.Datastore
	TemplateSet *isokit.TemplateSet
}
