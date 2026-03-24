package testdata

// String-based enum types.
type Severity string
type Status string

// Struct types.
type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Connection struct {
	Nodes    []*Node   `json:"nodes,omitempty"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type Node struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PageInfo struct {
	HasNextPage bool   `json:"hasNextPage"`
	EndCursor   string `json:"endCursor"`
}

// Interface type.
type Entity interface {
	GetID() string
}

// The main struct to flatten.
type Finding struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Score    *float64 `json:"score,omitempty"`
	Severity Severity `json:"severity"`
	Status   Status   `json:"status"`
	Deleted  bool     `json:"deleted"`
	Count    int64    `json:"count"`

	// Nested struct fields — should become map[string]any.
	Address    *Address    `json:"address,omitempty"`
	Tags       []*Tag      `json:"tags,omitempty"`
	Connection *Connection `json:"connection"`

	// Slice of scalars — should stay typed.
	Categories []string `json:"categories,omitempty"`

	// Interface field — should become map[string]any.
	Entity Entity `json:"entity"`

	// No JSON tag — should be skipped.
	Internal string

	// JSON tag "-" — should be skipped.
	Ignored string `json:"-"`
}

// A simple struct for testing multiple configs.
type Simple struct {
	ID    string `json:"id"`
	Value int64  `json:"value"`
}
