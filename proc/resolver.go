package proc

import (
	"fmt"

	common "github.com/ncabatoff/process-exporter"
)

type (
	// Resolver fills any additional fields in ProcAttributes
	Resolver interface {
		// Resolve fills any additional fields in ProcAttributes
		Resolve(*common.ProcAttributes, IDInfo)
		fmt.Stringer
	}
)
