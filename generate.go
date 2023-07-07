package jeppgen

import (
	_ "ariga.io/ogent"
)

//go:generate go run -mod=mod ./internal/ent/entc.go
//go:generate go run -mod=mod github.com/99designs/gqlgen
