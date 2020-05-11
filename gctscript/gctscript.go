package gctscript

import (
	"github.com/nbltrust/gocryptotrader/gctscript/modules"
	"github.com/nbltrust/gocryptotrader/gctscript/wrappers/gct"
)

// Setup configures the wrapper interface to use
func Setup() {
	modules.SetModuleWrapper(gct.Setup())
}
