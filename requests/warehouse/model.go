package warehouse

type WarehouseInfo struct {
	ID         uint64          `json:"warehouse_id,omitempty"`     // 仓库ID
	OutID      uint64          `json:"out_warehouse_id,omitempty"` // 外部仓库ID
	ShopID     uint64          `json:"shop_id,omitempty"`          // 店铺ID
	Name       string          `json:"name,omitempty"`             // 仓库名称
	Info       string          `json:"info,omitempty"`             // 仓库介绍
	UpdateTime int64           `json:"update_time,omitempty"`      // 更新时间
	CreateTime int64           `json:"create_time,omitempty"`      // 创建时间
	Addrs      []WarehouseAddr `json:"addr,omitempty"`             // 仓库地址数组
}

type WarehouseAddr struct {
	ID1        uint64 `json:"addr_id1,omitempty"`    // 一级地址
	ID2        uint64 `json:"addr_id2,omitempty"`    // 二级地址
	ID3        uint64 `json:"addr_id3,omitempty"`    // 三级地址
	ID4        uint64 `json:"addr_id4,omitempty"`    // 四级地址，目前只有三级地址，作为入参时可不传
	UpdateTime int64  `json:"update_time,omitempty"` // 更新时间，作为入参时可不传
	CreateTime int64  `json:"create_time,omitempty"` // 创建时间，作为入参时可不传
}
