package pkg

import (
	"github.com/slinkydeveloper/kfn/pkg/js"
)

type Bootstrapper interface {
	Bootstrap(functionName string, targetDirectory string) error
}

var (
	jsBootstrapper Bootstrapper = js.NewJSBootstrapper()
)

func ResolveBootstrapper(language Language) *Bootstrapper {
	switch language {
	case Javascript:
		return &jsBootstrapper
	default:
		return nil
	}
}
