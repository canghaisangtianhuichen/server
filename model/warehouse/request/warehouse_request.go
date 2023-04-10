package request

type Page struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}
type CustomersRequest struct {
	Name  string `json:"name" bind:"require"`
	Phone int    `json:"phone"`
	Email string `json:"email"`
}
type SuppliersRequest struct {
	Name  string `json:"name" bind:"require"`
	Phone int    `json:"phone"`
	Email string `json:"email"`
}
type WarehousesRequest struct {
	Name     string `json:"name" bind:"require"`
	Location string `json:"location" bind:"require"`
}
type AddWarehouseAndstaffRelations struct {
	Id          uint `json:"id" bind:"require"`
	WarehouseId uint `json:"warehouseId" bind:"require"`
}
type CancelWarehouseAndstaffRelations struct {
	Id uint `json:"id" bind:"require"`
}
type DeleteStaff struct {
	Id uint `json:"id" bind:"require"`
}
type DeleteCustomer struct {
	Id uint `json:"id" bind:"require"`
}
type DeleteSupplier struct {
	Id uint `json:"id" bind:"require"`
}
type DeleteWarehouse struct {
	Id uint `json:"id" bind:"require"`
}
type AddStaffRequest struct {
	Username     string `json:"userName" example:"用户名"`
	Password     string `json:"passWord" example:"密码"`
	NickName     string `json:"nickName" example:"昵称"`
	HeaderImg    string `json:"headerImg" example:"头像链接"`
	AuthorityId  uint   `json:"authorityId" swaggertype:"string" example:"int 角色id"`
	Enable       int    `json:"enable" swaggertype:"string" example:"int 是否启用"`
	AuthorityIds []uint `json:"authorityIds" swaggertype:"string" example:"[]uint 角色id"`
	Phone        string `json:"phone" example:"电话号码"`
	Email        string `json:"email" example:"电子邮箱"`
	WarehouseId  uint   `json:"warehouseId"`
}
type AddGoodsRequest struct {
	Name string `json:"name" bind:"require"`
}
type AddGoodsShelfRequest struct {
	Name      string `json:"name" bind:"require" `
	MaxWeight int    `json:"maxWeight" bind:"require,gt=0"` //大于零
}

type DeleteGoodRequest struct {
	Name        string `json:"name" bind:"require" `
	WarehouseId uint   `json:"warehouseId" bind:"require"`
}
type DeleteGoodsShelfRequest struct {
	Name        string `json:"name" bind:"require" `
	WarehouseId uint   `json:"warehouseId" bind:"require"`
}
type UpdateStaffRequest struct {
	Id    uint   `json:"id" bind:"require"`
	Name  string `json:"name" bind:"require"`
	Phone int    `json:"phone"`
	Email string `json:"email"`
}
type UpdateCustomerRequest struct {
	Id    uint   `json:"id" bind:"require"`
	Name  string `json:"name" bind:"require"`
	Phone int    `json:"phone"`
	Email string `json:"email"`
}
type UpdateSupplierRequest struct {
	Id    uint   `json:"id" bind:"require"`
	Name  string `json:"name" bind:"require"`
	Phone int    `json:"phone"`
	Email string `json:"email"`
}
type UpdateWarehouseRequest struct {
	Id       uint   `json:"id" bind:"require"`
	Name     string `json:"name" bind:"require"`
	Location string `json:"location" bind:"require"`
}
type InWarehouseRequest struct {
	TotalWeight int      `json:"totalWeight" bind:"require"`
	Type        string   `json:"type" bind:"require"` //0:调拨入库 1:采购入库
	FromId      uint     `json:"fromId" bind:"require"`
	GoodsId     []uint   `json:"goodsId" bind:"require"`
	Weight      []int    `json:"weight" bind:"require"`
	ShelfName   []string `json:"shelfName" bind:"require"`
}
type OutWarehouseRequest struct {
	TotalWeight int      `json:"totalWeight" bind:"require"`
	Type        string   `json:"type" bind:"require"` //0:调拨出库 1:采购出库
	ToId        uint     `json:"fromId" bind:"require"`
	GoodsId     []uint   `json:"goodsId" bind:"require"`
	Weight      []int    `json:"weight" bind:"require"`
	ShelfName   []string `json:"shelfName" bind:"require"`
}
