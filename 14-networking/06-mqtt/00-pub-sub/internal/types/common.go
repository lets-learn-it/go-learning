package types

type Entity struct {
	Type   string
	Facets map[string]Facet
}

type Facet map[string]any
