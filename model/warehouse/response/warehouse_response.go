package response

import "time"

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

type Base struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}
type WarehousesResponse struct {
	Base
	Name     string `json:"name"`
	Location string `json:"location"`
}
type InWarehousesResponse struct {
	Base
	OrderNumber string `json:"orderNumber"`
	WarehouseId uint   `json:"warehouseId"`
	StaffId     uint   `json:"staffId"`
	Weight      int    `json:"weight"`
	Type        string `json:"type"` //0:调拨入库 1:采购入库
	FromId      uint   `json:"fromId"`
}

type OutWarehousesResponse struct {
	Base
	OrderNumber string `json:"orderNumber"`
	WarehouseId uint   `json:"warehouseId"`
	StaffId     uint   `json:"staffId"`
	Weight      int    `json:"weight"`
	Type        string `json:"type"` //0:调拨出库 1:销售出库
	ToId        uint   `json:"toId"`
}

type InWarehousesDetailsResponse struct {
	Base
	OrderNumber string `json:"orderNumber"`
	WarehouseId uint   `json:"warehouseId"`
	StaffId     uint   `json:"staffId"`
	GoodsId     uint   `json:"goodsId"`
	Weight      int    `json:"weight"`
	ShelfName   string `json:"shelfName"`
}

type OutWarehousesDetailsResponse struct {
	Base
	OrderNumber string `json:"orderNumber"`
	WarehouseId uint   `json:"warehouseId"`
	StaffId     uint   `json:"staffId"`
	GoodsId     uint   `json:"goodsId"`
	Weight      int    `json:"weight"`
	ShelfName   string `json:"shelfName"`
}

type StaffsResponse struct {
	Base
	Name        string `json:"name"`
	Phone       int    `json:"phone"`
	Email       string `json:"email"`
	WarehouseId uint   `json:"warehouseId"`
}
type CustomersResponse struct {
	Base
	Name  string `json:"name"`
	Phone int    `json:"phone"`
	Email string `json:"email"`
}

type SupplierResponse struct {
	Base
	Name  string `json:"name"`
	Phone int    `json:"phone"`
	Email string `json:"email"`
}

type GoodsResponse struct {
	Base
	Name        string `json:"name"`
	WarehouseId uint   `json:"warehouseId"`
	Weight      int    `json:"weight"`
}
type GoodsShelfsResponse struct {
	Base
	Name           string `json:"name"`
	WarehouseId    uint   `json:"warehouseId"`
	GoodsId        uint   `json:"goodsId"`
	GoodsName      string `json:"goodsName"`
	RealTimeWeight int    `json:"realTimeWeight"`
	MaxWeight      int    `json:"maxWeight"`
}
