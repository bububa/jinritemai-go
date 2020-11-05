package address

type Province struct {
	ID   uint64 `json:"province_id,omitempty"` // long整型，省ID
	Name string `json:"province,omitempty"`    // 省名称，中文
	Pid  uint64 `json:"father_id,omitempty"`   // long整型，父节点ID
}

type City struct {
	ID   uint64 `json:"city_id,omitempty"`   // long整型，市ID
	Name string `json:"city,omitempty"`      // 市名称，中文
	Pid  uint64 `json:"father_id,omitempty"` // long整型，父节点ID
}

type Area struct {
	ID   uint64 `json:"area_id,omitempty"`   // long整型，区县ID
	Name string `json:"area,omitempty"`      // 区县名称，中文
	Pid  uint64 `json:"father_id,omitempty"` // long整型，父节点ID
}
