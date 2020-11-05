package category

type GoodsCategory struct {
	ID   uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CategoryProperty struct {
	ID           uint64           `json:"property_id,omitempty"`   // 属性id
	PropertyName string           `json:"property_name,omitempty"` // 属性名称
	Options      []PropertyOption `json:"options,omitempty"`       // 属性可选值列表，为空时需要手动填写
	Required     bool             `json:"required,omitempty"`      // true：创建商品时该属性字段必填; false：创建商品时该属性字段选填
}

type PropertyOption struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}
