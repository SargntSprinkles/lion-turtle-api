package server

import (
	_ "embed"

	"github.com/SargntSprinkles/lion-turtle-api/server/resolvers/moves"
	"github.com/SargntSprinkles/lion-turtle-api/server/resolvers/playbooks"
	"github.com/SargntSprinkles/lion-turtle-api/server/resolvers/techniques"
)

//go:embed schema.graphql
var schema string

type Resolver struct {
	playbooks.PlaybookResolver
	techniques.TechniqueResolver
	moves.MoveResolver
}

func (r *Resolver) schema() string {
	return schema
}
