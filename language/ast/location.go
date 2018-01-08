package ast

import (
	"github.com/graphql-go/graphql/language/source"
)

type Location struct {
	Start  int
	End    int
	Source *source.Source
}

func NewLocation(loc Location) Location {
	return Location{
		Start:  loc.Start,
		End:    loc.End,
		Source: loc.Source,
	}
}
