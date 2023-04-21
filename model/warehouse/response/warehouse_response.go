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
	UpdatedAt time.Time `json:"updatedAt"`
}
type WarehousesResponse struct {
	Base
	Name     string `json:"name"`
	Location string `json:"location"`
}
type InWarehousesResponse struct {
	Base
	OrderNumber   string `json:"orderNumber"`
	WarehouseId   uint   `json:"warehouseId"`
	WarehouseName string `json:"warehouseName"`
	StaffId       uint   `json:"staffId"`
	StaffName     string `json:"staffName"`
	Weight        int    `json:"weight"`
	Type          string `json:"type"` //0:调拨入库 1:采购入库
	FromId        uint   `json:"fromId"`
	FromWhere     string `json:"fromWhere"`
}

type OutWarehousesResponse struct {
	Base
	OrderNumber   string `json:"orderNumber"`
	WarehouseId   uint   `json:"warehouseId"`
	WarehouseName string `json:"warehouseName"`
	StaffId       uint   `json:"staffId"`
	StaffName     string `json:"staffName"`
	Weight        int    `json:"weight"`
	Type          string `json:"type"` //0:调拨出库 1:销售出库
	ToId          uint   `json:"toId"`
	ToWhere       string `json:"toWhere"`
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
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	WarehouseId   uint   `json:"warehouseId"`
	WarehouseName string `json:"warehouseName"`
	Status        string `json:"status"`
}
type CustomersResponse struct {
	Base
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type SupplierResponse struct {
	Base
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type GoodsResponse struct {
	Base
	Name          string `json:"name"`
	WarehouseId   uint   `json:"warehouseId"`
	WarehouseName string `json:"warehouseName"`
	Weight        int    `json:"weight"`
}
type GoodsShelfsResponse struct {
	Base
	Name           string `json:"name"`
	WarehouseId    uint   `json:"warehouseId"`
	WarehouseName  string `json:"warehouseName"`
	GoodsId        uint   `json:"goodsId"`
	GoodsName      string `json:"goodsName"`
	RealTimeWeight int    `json:"realTimeWeight"`
	MaxWeight      int    `json:"maxWeight"`
}
