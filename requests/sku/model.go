package sku

type SkuDetail struct {
	ID              uint64         `json:"id,omitempty"`
	OutID           uint64         `json:"out_sku_id,omitempty"`
	ProductID       uint64         `json:"product_id,omitempty"`
	ProductIDStr    string         `json:"product_id_str,omitempty"`
	SpecID          uint64         `json:"spec_id,omitempty"`
	SpecDetailID1   uint64         `json:"spec_detail_id1,omitempty"`
	SpecDetailID2   uint64         `json:"spec_detail_id2,omitempty"`
	SpecDetailID3   uint64         `json:"spec_detail_id3,omitempty"`
	SpecDetailName1 string         `json:"spec_detail_name1,omitempty"`
	SpecDetailName2 string         `json:"spec_detail_name2,omitempty"`
	SpecDetailName3 string         `json:"spec_detail_name3,omitempty"`
	StockNum        int            `json:"stock_num,omitempty"`
	Price           int            `json:"price,omitempty"`
	SettlementPrice int            `json:"settlement_price,omitempty"`
	Code            string         `json:"code,omitempty"`
	CreateTime      int64          `json:"create_time,omitempty"`
	StockType       int            `json:"sku_type,omitempty"`  // 库存类型：0普通库存，1区域库存
	StockMap        map[string]int `json:"stock_map,omitempty"` // 区域仓库存信息，out_warehouse_id与stock_num对应关系
}

type SkuStock struct {
	SkuID    uint64         `json:"sku_id,omitempty"`
	OutID    uint64         `json:"out_sku_id,omitempty"`
	SkuType  uint           `json:"sku_type,omitempty"`  // sku库存类型，0表示普通，1表示区域仓库存
	StockNum int            `json:"stock_num,omitempty"` // 如果sku_type=0，表示全国库存; 如果sku_type=1，且入参有out_warehouse_id，则表示该仓库的库存; 如果sku_type=1，且入参无out_warehouse_id，则为空
	StockMap map[string]int `json:"stock_map,omitempty"` // 如果sku_type=0，为空如果sku_type=1，则为区域仓库存映射表，key为out_warehouse_id，value为库存
}
