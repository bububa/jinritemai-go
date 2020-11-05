package spec

type SpecDetail struct {
	ID    uint64 `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Specs []Spec `json:"specs,omitempty"`
}

type Spec struct {
	ID     uint64       `json:"id,omitempty"`
	SpecID uint64       `json:"spec_id,omitempty"`
	Name   string       `json:"name,omitempty"`
	Pid    uint64       `json:"pid,omitempty"`
	IsLeaf uint         `json:"is_leaf,omitempty"`
	Values []SpecValues `json:"values,omitempty"`
}

type SpecValues struct {
	ID     uint64 `json:"id,omitempty"`
	SpecID uint64 `json:"spec_id,omitempty"`
	Name   string `json:"name,omitempty"`
	Pid    uint64 `json:"pid,omitempty"`
	IsLeaf uint   `json:"is_leaf,omitempty"`
	Status uint   `json:"status,omitempty"`
}
