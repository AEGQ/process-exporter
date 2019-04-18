package common

import "fmt"

type (
	ProcAttributes struct {
		Pid      int
		PPid     int
		Name     string
		Cmdline  []string
		Username string
		Pod      string
	}

	// Resolver fills any additional fields in ProcAttributes
	Resolver interface {
		// Resolve fills any additional fields in ProcAttributes
		Resolve(*ProcAttributes)
		fmt.Stringer
	}

	MatchNamer interface {
		// MatchAndName returns false if the match failed, otherwise
		// true and the resulting name.
		MatchAndName(ProcAttributes) (bool, string)
		AddResolver(Resolver)
		fmt.Stringer
	}
)
