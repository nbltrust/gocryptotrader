package wrappers

import (
	"github.com/nbltrust/gocryptotrader/gctscript/modules"
	"github.com/nbltrust/gocryptotrader/gctscript/wrappers/validator"
)

// GetWrapper returns the instance of each wrapper to use
func GetWrapper() modules.GCT {
	if validator.IsTestExecution.Load() == true {
		return validator.Wrapper{}
	}
	return modules.Wrapper
}
