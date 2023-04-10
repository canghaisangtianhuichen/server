package warehouse

import (
	"gorm.io/gorm"
	"time"
)

type Base struct {
	Id        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
type Warehouses struct {
	Base
	Name     string `json:"name"`
	Location string `json:"location"`
}
type InWarehouses struct {
	Base
	OrderNumber string `json:"orderNumber"`
	WarehouseId uint   `json:"warehouseId"`
	StaffId     uint   `json:"staffId"`
	Weight      int    `json:"weight"`
	Type        string `json:"type"` //0:调拨入库 1:采购入库
	FromId      uint   `json:"fromId"`
}

type OutWarehouses struct {
	Base
	OrderNumber string `json:"orderNumber"`
	WarehouseId uint   `json:"warehouseId"`
	StaffId     uint   `json:"staffId"`
	Weight      int    `json:"weight"`
	Type        string `json:"type"` //0:调拨出库 1:销售出库
	ToId        uint   `json:"toId"`
}

type InWarehousesDetails struct {
	Base
	OrderNumber string `json:"orderNumber"`
	WarehouseId uint   `json:"warehouseId"`
	StaffId     uint   `json:"staffId"`
	GoodsId     uint   `json:"goodsId"`
	Weight      int    `json:"weight"`
	ShelfName   string `json:"shelfName"`
}

type OutWarehousesDetails struct {
	Base
	OrderNumber string `json:"orderNumber"`
	WarehouseId uint   `json:"warehouseId"`
	StaffId     uint   `json:"staffId"`
	GoodsId     uint   `json:"goodsId"`
	Weight      int    `json:"weight"`
	ShelfName   string `json:"shelfName"`
}

type Staffs struct {
	Base
	SysId       uint   `json:"sysId"`
	Name        string `json:"name"`
	Phone       int    `json:"phone"`
	Email       string `json:"email"`
	WarehouseId uint   `json:"warehouseId"`
}
type Customers struct {
	Base
	Name  string `json:"name"`
	Phone int    `json:"phone"`
	Email string `json:"email"`
}

type Suppliers struct {
	Base
	Name  string `json:"name"`
	Phone int    `json:"phone"`
	Email string `json:"email"`
}

type Goods struct {
	Base
	Name        string `json:"name"`
	WarehouseId uint   `json:"warehouseId"`
	Weight      int    `json:"weight"`
}
type GoodsShelfs struct {
	Base
	Name           string `json:"name"`
	WarehouseId    uint   `json:"warehouseId"`
	RealTimeWeight int    `json:"realTimeWeight"`
	MaxWeight      int    `json:"maxWeight"`
}
