package system

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/warehouse"
	"github.com/flipped-aurora/gin-vue-admin/server/model/warehouse/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/warehouse/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type WarehouseService struct{}

func (warehouseService *WarehouseService) GetV1WarehousesList(pageInfo request.Page) (data []response.WarehousesResponse, total int64, err error) {
	sql := "select * from warehouses where deleted_at is null limit ?,?"
	err1 := global.GVA_DB.Raw(sql, (pageInfo.Page-1)*pageInfo.PageSize, pageInfo.PageSize).Scan(&data)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	sql = "select count(id) from warehouses where deleted_at is null "
	err1 = global.GVA_DB.Raw(sql).Scan(&total)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	return data, total, nil
}
func (warehouseService *WarehouseService) GetV1GoodsList(pageInfo request.Page) (data []response.GoodsResponse, total int64, err error) {
	sql := "select * from goods where deleted_at is null limit ?,?"
	err1 := global.GVA_DB.Raw(sql, (pageInfo.Page-1)*pageInfo.PageSize, pageInfo.PageSize).Scan(&data)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	sql = " select count(id) from goods where deleted_at is null"
	err1 = global.GVA_DB.Raw(sql).Scan(&total)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	return data, total, nil
}
func (warehouseService *WarehouseService) GetV1StaffsList(pageInfo request.Page) (data []response.StaffsResponse, total int64, err error) {
	sql := "select * from staffs where deleted_at is null limit ?,?"
	err1 := global.GVA_DB.Raw(sql, (pageInfo.Page-1)*pageInfo.PageSize, pageInfo.PageSize).Scan(&data)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	sql = " select count(id) from staffs where deleted_at is null"
	err1 = global.GVA_DB.Raw(sql).Scan(&total)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	return data, total, nil
}
func (warehouseService *WarehouseService) GetV1CustomersList(pageInfo request.Page) (data []response.CustomersResponse, total int64, err error) {
	sql := "select * from customers where deleted_at is null limit ?,?"
	err1 := global.GVA_DB.Raw(sql, (pageInfo.Page-1)*pageInfo.PageSize, pageInfo.PageSize).Scan(&data)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	sql = " select count(id) from customers where deleted_at is null"
	err1 = global.GVA_DB.Raw(sql).Scan(&total)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	return data, total, nil
}
func (warehouseService *WarehouseService) GetV1SuppliersList(pageInfo request.Page) (data []response.SupplierResponse, total int64, err error) {
	sql := "select * from suppliers where deleted_at is null limit ?,?"
	err1 := global.GVA_DB.Raw(sql, (pageInfo.Page-1)*pageInfo.PageSize, pageInfo.PageSize).Scan(&data)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	sql = " select count(id) from suppliers where deleted_at is null"
	err1 = global.GVA_DB.Raw(sql).Scan(&total)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	return data, total, nil
}
func (warehouseService *WarehouseService) GetV1GoodsShelfsList(pageInfo request.Page) (data []response.GoodsShelfsResponse, total int64, err error) {
	sql := "select gs.id,gs.`name`,gs.warehouse_id,gs.goods_id,g.name GoodsName,gs.real_time_weight,gs.max_weight,gs.created_at,gs.updated_at from goods_shelfs gs left join goods g on (gs.goods_id=g.id and gs.warehouse_id=g.warehouse_id ) where gs.deleted_at is null and g.deleted_at is null limit ?,?"
	err1 := global.GVA_DB.Raw(sql, (pageInfo.Page-1)*pageInfo.PageSize, pageInfo.PageSize).Scan(&data)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	sql = " select count(id) from goods_shelfs where deleted_at is null"
	err1 = global.GVA_DB.Raw(sql).Scan(&total)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	return data, total, nil
}
func (warehouseService *WarehouseService) GetV1OutWarehousesList(pageInfo request.Page) (data []response.OutWarehousesResponse, total int64, err error) {
	sql := "select * from out_warehouses where deleted_at is null limit ?,?"
	err1 := global.GVA_DB.Raw(sql, (pageInfo.Page-1)*pageInfo.PageSize, pageInfo.PageSize).Scan(&data)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	sql = " select count(id) from out_warehouses where deleted_at is null"
	err1 = global.GVA_DB.Raw(sql).Scan(&total)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	return data, total, nil
}
func (warehouseService *WarehouseService) GetV1OutWarehousesDetail(orderNumber string, pageInfo request.Page) (data []response.OutWarehousesDetailsResponse, total int64, err error) {
	sql := "select * from out_warehouses_details where deleted_at is null and order_number=? limit ?,?"
	err1 := global.GVA_DB.Raw(sql, orderNumber, (pageInfo.Page-1)*pageInfo.PageSize, pageInfo.PageSize).Scan(&data)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	sql = " select count(id) from out_warehouses_details where deleted_at is null and order_number=?"
	err1 = global.GVA_DB.Raw(sql, orderNumber).Scan(&total)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	return data, total, nil
}
func (warehouseService *WarehouseService) GetV1InWarehousesList(pageInfo request.Page) (data []response.InWarehousesResponse, total int64, err error) {
	sql := "select * from in_warehouses where deleted_at is null limit ?,?"
	err1 := global.GVA_DB.Raw(sql, (pageInfo.Page-1)*pageInfo.PageSize, pageInfo.PageSize).Scan(&data)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	sql = " select count(id) from in_warehouses where deleted_at is null"
	err1 = global.GVA_DB.Raw(sql).Scan(&total)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	return data, total, nil
}
func (warehouseService *WarehouseService) GetV1InWarehousesDetail(orderNumber string, pageInfo request.Page) (data []response.InWarehousesDetailsResponse, total int64, err error) {
	sql := "select * from in_warehouses_details where deleted_at is null and order_number=? limit ?,?"
	err1 := global.GVA_DB.Raw(sql, orderNumber, (pageInfo.Page-1)*pageInfo.PageSize, pageInfo.PageSize).Scan(&data)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	sql = " select count(id) from in_warehouses_details where deleted_at is null and order_number=?"
	err1 = global.GVA_DB.Raw(sql, orderNumber).Scan(&total)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	return data, total, nil
}

func (warehouseService *WarehouseService) AddCustomer(info request.CustomersRequest) (err error) {
	var total int64
	var customer warehouse.Customers
	if info.Phone != 0 {
		sql := "select count(id) from customers where deleted_at is null and phone=? "
		err1 := global.GVA_DB.Raw(sql, info.Phone).Scan(&total)
		if err1.Error != nil {
			return errors.New("系统错误")
		}
		if total != 0 {
			return errors.New("该手机号已有客户绑定")
		}
	}
	if info.Email != "" {
		sql := "select count(id) from customers where deleted_at is null and email=? "
		err1 := global.GVA_DB.Raw(sql, info.Email).Scan(&total)
		if err1.Error != nil {
			return errors.New("系统错误")
		}
		if total != 0 {
			return errors.New("该邮箱已有客户绑定")
		}
	}
	customer.Name = info.Name
	customer.Phone = info.Phone
	customer.Email = info.Email
	customer.CreatedAt = time.Now()
	err1 := global.GVA_DB.Table("customers").Create(&customer)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	return nil
}
func (warehouseService *WarehouseService) AddSupplier(info request.SuppliersRequest) (err error) {
	var total int64
	var supplier warehouse.Suppliers
	if info.Phone != 0 {
		sql := "select count(id) from suppliers where deleted_at is null and phone=? "
		err1 := global.GVA_DB.Raw(sql, info.Phone).Scan(&total)
		if err1.Error != nil {
			return errors.New("系统错误")
		}
		if total != 0 {
			return errors.New("该手机号已有客户绑定")
		}
	}
	if info.Email != "" {
		sql := "select count(id) from suppliers where deleted_at is null and email=? "
		err1 := global.GVA_DB.Raw(sql, info.Email).Scan(&total)
		if err1.Error != nil {
			return errors.New("系统错误")
		}
		if total != 0 {
			return errors.New("该邮箱已有客户绑定")
		}
	}
	supplier.Name = info.Name
	supplier.Phone = info.Phone
	supplier.Email = info.Email
	supplier.CreatedAt = time.Now()
	err1 := global.GVA_DB.Table("suppliers").Create(&supplier)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	return nil
}
func (warehouseService *WarehouseService) AddWarehouse(info request.WarehousesRequest) (err error) {
	var total int64
	var warehouse warehouse.Warehouses
	sql := "select count(id) from warehouses where deleted_at is null and `name`=?  "
	err1 := global.GVA_DB.Raw(sql, info.Name).Scan(&total)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	if total != 0 {
		return errors.New("该仓库名已存在")
	}

	sql = "select count(id) from warehouses where deleted_at is null and  location=? "
	err1 = global.GVA_DB.Raw(sql, info.Location).Scan(&total)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	if total != 0 {
		return errors.New("该地址已有仓库")
	}
	warehouse.Name = info.Name
	warehouse.Location = info.Location
	warehouse.CreatedAt = time.Now()
	err1 = global.GVA_DB.Table("warehouses").Create(&warehouse)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	return nil
}
func (warehouseService *WarehouseService) AddWarehouseAndstaffRelations(info request.AddWarehouseAndstaffRelations) (err error) {
	var warehouseId uint
	sql := "select warehouse_id from staffs where deleted_at is null and id=?  "
	err1 := global.GVA_DB.Raw(sql, info.Id).Scan(&warehouseId)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	if warehouseId != 0 {
		return errors.New("该员工已分配到其他仓库")
	}

	err1 = global.GVA_DB.Table("staffs").Where("id=?", info.Id).Update("warehouse_id", info.WarehouseId)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	return nil
}
func (warehouseService *WarehouseService) CancelWarehouseAndstaffRelations(info request.CancelWarehouseAndstaffRelations) (err error) {
	var warehouseId uint
	sql := "select warehouse_id from staffs where deleted_at is null and id=?  "
	err1 := global.GVA_DB.Raw(sql, info.Id).Scan(&warehouseId)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	if warehouseId == 0 {
		return errors.New("该员工未分配到仓库")
	}

	err1 = global.GVA_DB.Table("staffs").Where("id=?", info.Id).Update("warehouse_id", 0)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	return nil
}
func (warehouseService *WarehouseService) AddStaff(u system.SysUser, warehouseId uint) (userInter system.SysUser, err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var user system.SysUser
		var staff warehouse.Staffs
		if !errors.Is(tx.Where("username = ? and deleted_at is null", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
			return errors.New("用户名已注册")
		}
		// 否则 附加uuid 密码hash加密 注册
		u.Password = utils.BcryptHash(u.Password)
		u.UUID = uuid.NewV4()
		err = tx.Create(&u).Error
		if err != nil {
			return err
		}
		var temUser system.SysUser
		//在员工表增加数据
		err1 := tx.Table("sys_users").Where("username = ? and deleted_at is null", u.Username).Scan(&temUser)
		if err1.Error != nil {
			return err1.Error
		}
		staff.Name = u.NickName
		staff.WarehouseId = warehouseId
		staff.Phone, _ = strconv.Atoi(u.Phone)
		staff.Email = u.Email
		staff.SysId = temUser.ID
		err1 = tx.Table("staffs").Create(&staff)
		if err1.Error != nil {
			return err1.Error
		}
		return nil
	})
	return u, err
}
func (warehouseService *WarehouseService) DeleteStaff(info request.DeleteStaff) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var total int64
		var sysId uint
		sql := "select count(id) from staffs where deleted_at is null and id=?  "
		err1 := global.GVA_DB.Raw(sql, info.Id).Scan(&total)
		if err1.Error != nil {
			return errors.New("系统错误")
		}
		if total == 0 {
			return errors.New("非法请求")
		}
		sql = "select sys_id from staffs where deleted_at is null and id=?  "
		err1 = global.GVA_DB.Raw(sql, info.Id).Scan(&sysId)
		if err1.Error != nil {
			return errors.New("系统错误")
		}
		err1 = global.GVA_DB.Table("staffs").Where("id=?", info.Id).Update("deleted_at", time.Now())
		if err1.Error != nil {
			return errors.New("系统错误")
		}
		err1 = global.GVA_DB.Table("sys_users").Where("id=?", sysId).Update("deleted_at", time.Now())
		if err1.Error != nil {
			return errors.New("系统错误")
		}
		return nil
	})
	return err
}
func (warehouseService *WarehouseService) DeleteCustomer(info request.DeleteCustomer) (err error) {
	var total int64
	sql := "select count(id) from customers where deleted_at is null and id=?  "
	err1 := global.GVA_DB.Raw(sql, info.Id).Scan(&total)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	if total == 0 {
		return errors.New("非法请求")
	}

	err1 = global.GVA_DB.Table("customers").Where("id=?", info.Id).Update("deleted_at", time.Now())
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	return nil
}
func (warehouseService *WarehouseService) DeleteSupplier(info request.DeleteSupplier) (err error) {
	var total int64
	sql := "select count(id) from suppliers where deleted_at is null and id=?  "
	err1 := global.GVA_DB.Raw(sql, info.Id).Scan(&total)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	if total == 0 {
		return errors.New("非法请求")
	}

	err1 = global.GVA_DB.Table("suppliers").Where("id=?", info.Id).Update("deleted_at", time.Now())
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	return nil
}
func (warehouseService *WarehouseService) DeleteWarehouse(info request.DeleteWarehouse) (err error) {
	var total int64
	var weight int
	sql := "select count(id) from warehouses where deleted_at is null and id=?  "
	err1 := global.GVA_DB.Raw(sql, info.Id).Scan(&total)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	if total == 0 {
		return errors.New("非法请求")
	}
	sql = "select ifnull(sum(weight),0) from goods where deleted_at is null and warehouse_id=? "
	err1 = global.GVA_DB.Raw(sql, info.Id).Scan(&weight)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	if weight != 0 {
		return errors.New("删除仓库失败！请先将仓库货物清空")
	}
	err1 = global.GVA_DB.Table("warehouses").Where("id=?", info.Id).Update("deleted_at", time.Now())
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	return nil
}
func (warehouseService *WarehouseService) UpdateStaff(info request.UpdateStaffRequest) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var total int64
		var sysId uint
		sql := "select count(id) from staffs where deleted_at is null and id=?  "
		err1 := global.GVA_DB.Raw(sql, info.Id).Scan(&total)
		if err1.Error != nil {
			return errors.New("系统错误")
		}
		if total == 0 {
			return errors.New("非法请求")
		}
		sql = "select sys_id from staffs where deleted_at is null and id=?  "
		err1 = global.GVA_DB.Raw(sql, info.Id).Scan(&sysId)
		if err1.Error != nil {
			return errors.New("系统错误")
		}
		err1 = global.GVA_DB.Table("staffs").Where("id=? and deleted_at is null", info.Id).Update("name", info.Name).Update("phone", info.Phone).Update("email", info.Email)
		if err1.Error != nil {
			return errors.New("系统错误")
		}
		err1 = global.GVA_DB.Table("sys_users").Where("id=? and deleted_at is null", sysId).Update("phone", info.Phone).Update("email", info.Email)
		if err1.Error != nil {
			return errors.New("系统错误")
		}
		return nil
	})
	return err
}
func (warehouseService *WarehouseService) UpdateCustomer(info request.UpdateCustomerRequest) (err error) {
	var total int64
	sql := "select count(id) from customers where deleted_at is null and id=?  "
	err1 := global.GVA_DB.Raw(sql, info.Id).Scan(&total)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	if total == 0 {
		return errors.New("非法请求")
	}
	err1 = global.GVA_DB.Table("customers").Where("id=? and deleted_at is null", info.Id).Update("name", info.Name).Update("phone", info.Phone).Update("email", info.Email)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	return nil
}
func (warehouseService *WarehouseService) UpdateSupplier(info request.UpdateSupplierRequest) (err error) {
	var total int64
	sql := "select count(id) from suppliers where deleted_at is null and id=?  "
	err1 := global.GVA_DB.Raw(sql, info.Id).Scan(&total)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	if total == 0 {
		return errors.New("非法请求")
	}
	err1 = global.GVA_DB.Table("suppliers").Where("id=? and deleted_at is null", info.Id).Update("name", info.Name).Update("phone", info.Phone).Update("email", info.Email)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	return nil
}
func (warehouseService *WarehouseService) UpdateWarehouse(info request.UpdateWarehouseRequest) (err error) {
	var total int64
	sql := "select count(id) from warehouses where deleted_at is null and id=?  "
	err1 := global.GVA_DB.Raw(sql, info.Id).Scan(&total)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	if total == 0 {
		return errors.New("非法请求")
	}
	err1 = global.GVA_DB.Table("warehouses").Where("id=? and deleted_at is null", info.Id).Update("name", info.Name).Update("location", info.Location)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	return nil
}

//------------------------------------v2---------------------------------------------

func (warehouseService *WarehouseService) GetV2GoodsList(sysId uint, pageInfo request.Page) (data []response.GoodsResponse, total int64, err error) {
	var warehouseId uint
	sql := "select warehouse_id from staffs where deleted_at is null and sys_id=? "
	err1 := global.GVA_DB.Raw(sql, sysId).Scan(&warehouseId)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}

	sql = "select * from goods where deleted_at is null and warehouse_id=? limit ?,?"
	err1 = global.GVA_DB.Raw(sql, warehouseId, (pageInfo.Page-1)*pageInfo.PageSize, pageInfo.PageSize).Scan(&data)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	sql = " select count(id) from goods where deleted_at is null and warehouse_id=?"
	err1 = global.GVA_DB.Raw(sql, warehouseId).Scan(&total)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	return data, total, nil
}
func (warehouseService *WarehouseService) GetV2StaffsList(sysId uint, pageInfo request.Page) (data []response.StaffsResponse, total int64, err error) {
	var warehouseId uint
	sql := "select warehouse_id from staffs where deleted_at is null and sys_id=? "
	err1 := global.GVA_DB.Raw(sql, sysId).Scan(&warehouseId)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	sql = "select * from staffs where deleted_at is null and warehouse_id=? limit ?,?"
	err1 = global.GVA_DB.Raw(sql, warehouseId, (pageInfo.Page-1)*pageInfo.PageSize, pageInfo.PageSize).Scan(&data)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	sql = "select count(id) from staffs where deleted_at is null and warehouse_id=? "
	err1 = global.GVA_DB.Raw(sql, warehouseId).Scan(&total)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	return data, total, nil
}
func (warehouseService *WarehouseService) GetV2CustomersList(pageInfo request.Page) (data []response.CustomersResponse, total int64, err error) {
	sql := "select * from customers where deleted_at is null limit ?,?"
	err1 := global.GVA_DB.Raw(sql, (pageInfo.Page-1)*pageInfo.PageSize, pageInfo.PageSize).Scan(&data)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	sql = " select count(id) from customers where deleted_at is null"
	err1 = global.GVA_DB.Raw(sql).Scan(&total)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	return data, total, nil
}
func (warehouseService *WarehouseService) GetV2SuppliersList(pageInfo request.Page) (data []response.SupplierResponse, total int64, err error) {
	sql := "select * from suppliers where deleted_at is null limit ?,?"
	err1 := global.GVA_DB.Raw(sql, (pageInfo.Page-1)*pageInfo.PageSize, pageInfo.PageSize).Scan(&data)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	sql = " select count(id) from suppliers where deleted_at is null"
	err1 = global.GVA_DB.Raw(sql).Scan(&total)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	return data, total, nil
}
func (warehouseService *WarehouseService) GetV2GoodsShelfsList(sysId uint, pageInfo request.Page) (data []response.GoodsShelfsResponse, total int64, err error) {
	var warehouseId uint
	sql := "select warehouse_id from staffs where deleted_at is null and sys_id=?"
	err1 := global.GVA_DB.Raw(sql, sysId).Scan(&warehouseId)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	sql = "select gs.id,gs.`name`,gs.warehouse_id ,gs.goods_id,g.name GoodsName,gs.real_time_weight,gs.max_weight,gs.created_at,gs.updated_at from goods_shelfs gs left join goods g on (gs.goods_id=g.id and gs.warehouse_id=g.warehouse_id ) where gs.deleted_at is null and g.deleted_at is null and gs.warehouse_id=? limit ?,?"
	err1 = global.GVA_DB.Raw(sql, warehouseId, (pageInfo.Page-1)*pageInfo.PageSize, pageInfo.PageSize).Scan(&data)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	sql = " select count(id) from goods_shelfs where deleted_at is null and warehouse_id=?"
	err1 = global.GVA_DB.Raw(sql, warehouseId).Scan(&total)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	return data, total, nil
}
func (warehouseService *WarehouseService) GetV2OutWarehousesList(sysId uint, pageInfo request.Page) (data []response.OutWarehousesResponse, total int64, err error) {
	var warehouseId uint
	sql := "select warehouse_id from staffs where deleted_at is null and sys_id=?"
	err1 := global.GVA_DB.Raw(sql, sysId).Scan(&warehouseId)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	sql = "select * from out_warehouses where deleted_at is null and warehouse_id=? limit ?,?"
	err1 = global.GVA_DB.Raw(sql, warehouseId, (pageInfo.Page-1)*pageInfo.PageSize, pageInfo.PageSize).Scan(&data)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	sql = " select count(id) from out_warehouses where deleted_at is null and warehouse_id=?"
	err1 = global.GVA_DB.Raw(sql, warehouseId).Scan(&total)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	return data, total, nil
}
func (warehouseService *WarehouseService) GetV2OutWarehousesDetail(orderNumber string, sysId uint, pageInfo request.Page) (data []response.OutWarehousesDetailsResponse, total int64, err error) {
	var warehouseId uint
	sql := "select warehouse_id from staffs where deleted_at is null and sys_id=?"
	err1 := global.GVA_DB.Raw(sql, sysId).Scan(&warehouseId)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	sql = "select * from out_warehouses_details where deleted_at is null and order_number=? and warehouse_id=? limit ?,?"
	err1 = global.GVA_DB.Raw(sql, orderNumber, warehouseId, (pageInfo.Page-1)*pageInfo.PageSize, pageInfo.PageSize).Scan(&data)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	sql = " select count(id) from out_warehouses_details where deleted_at is null and order_number=? and warehouse_id=?"
	err1 = global.GVA_DB.Raw(sql, orderNumber, warehouseId).Scan(&total)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	return data, total, nil
}
func (warehouseService *WarehouseService) GetV2InWarehousesList(sysId uint, pageInfo request.Page) (data []response.InWarehousesResponse, total int64, err error) {
	var warehouseId uint
	sql := "select warehouse_id from staffs where deleted_at is null and sys_id=?"
	err1 := global.GVA_DB.Raw(sql, sysId).Scan(&warehouseId)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	sql = "select * from in_warehouses where deleted_at is null and warehouse_id=? limit ?,?"
	err1 = global.GVA_DB.Raw(sql, warehouseId, (pageInfo.Page-1)*pageInfo.PageSize, pageInfo.PageSize).Scan(&data)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	sql = " select count(id) from in_warehouses where deleted_at is null and warehouse_id=?"
	err1 = global.GVA_DB.Raw(sql, warehouseId).Scan(&total)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	return data, total, nil
}
func (warehouseService *WarehouseService) GetV2InWarehousesDetail(orderNumber string, sysId uint, pageInfo request.Page) (data []response.InWarehousesDetailsResponse, total int64, err error) {
	var warehouseId uint
	sql := "select warehouse_id from staffs where deleted_at is null and sys_id=?"
	err1 := global.GVA_DB.Raw(sql, sysId).Scan(&warehouseId)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	sql = "select * from in_warehouses_details where deleted_at is null and order_number=? and warehouse_id=? limit ?,?"
	err1 = global.GVA_DB.Raw(sql, orderNumber, warehouseId, (pageInfo.Page-1)*pageInfo.PageSize, pageInfo.PageSize).Scan(&data)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	sql = " select count(id) from in_warehouses_details where deleted_at is null and order_number=? and warehouse_id=?"
	err1 = global.GVA_DB.Raw(sql, orderNumber, warehouseId).Scan(&total)
	if err1.Error != nil {
		return data, total, errors.New("系统错误")
	}
	return data, total, nil
}

func (warehouseService *WarehouseService) AddGood(info request.AddGoodsRequest, sysId uint) (err error) {
	var total int64
	var goods warehouse.Goods
	var warehouseId uint
	sql := "select warehouse_id from staffs where deleted_at is null and sys_id=?"
	err1 := global.GVA_DB.Raw(sql, sysId).Scan(&warehouseId)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	sql = "select count(id) from goods where deleted_at is null and `name`=? and warehouse_id=?  "
	err1 = global.GVA_DB.Raw(sql, info.Name, warehouseId).Scan(&total)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	if total != 0 {
		return errors.New("仓库已存在该货物")
	}

	goods.Name = info.Name
	goods.WarehouseId = warehouseId
	goods.CreatedAt = time.Now()
	err1 = global.GVA_DB.Table("goods").Create(&goods)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	return nil
}
func (warehouseService *WarehouseService) AddGoodsShelf(info request.AddGoodsShelfRequest, sysId uint) (err error) {
	var total int64
	var goodsShelf warehouse.GoodsShelfs
	var warehouseId uint
	sql := "select warehouse_id from staffs where deleted_at is null and sys_id=?"
	err1 := global.GVA_DB.Raw(sql, sysId).Scan(&warehouseId)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	sql = "select count(id) from goods_shelfs where deleted_at is null and `name`=? and warehouse_id=?  "
	err1 = global.GVA_DB.Raw(sql, info.Name, warehouseId).Scan(&total)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	if total != 0 {
		return errors.New("仓库已存在该货架")
	}

	goodsShelf.Name = info.Name
	goodsShelf.MaxWeight = info.MaxWeight
	goodsShelf.WarehouseId = warehouseId
	goodsShelf.CreatedAt = time.Now()
	err1 = global.GVA_DB.Table("goods_shelfs").Create(&goodsShelf)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	return nil
}
func (warehouseService *WarehouseService) DeleteGood(info request.DeleteGoodRequest, sysId uint) (err error) {
	var total int64
	var warehouseId uint
	sql := "select warehouse_id from staffs where deleted_at is null and sys_id=?"
	err1 := global.GVA_DB.Raw(sql, sysId).Scan(&warehouseId)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	if info.WarehouseId != warehouseId {
		return errors.New("非法请求")
	}
	sql = "select count(id) from goods where deleted_at is null and `name`=? and warehouse_id=?  "
	err1 = global.GVA_DB.Raw(sql, info.Name, warehouseId).Scan(&total)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	if total != 1 {
		return errors.New("仓库不存在该货物")
	}
	err1 = global.GVA_DB.Table("goods").Where("name=? and warehouse_id=? and deleted_at is null", info.Name, info.WarehouseId).Update("deleted_at", time.Now())
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	return nil
}
func (warehouseService *WarehouseService) DeleteGoodsShelf(info request.DeleteGoodsShelfRequest, sysId uint) (err error) {
	var total int64
	var warehouseId uint
	sql := "select warehouse_id from staffs where deleted_at is null and sys_id=?"
	err1 := global.GVA_DB.Raw(sql, sysId).Scan(&warehouseId)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	if info.WarehouseId != warehouseId {
		return errors.New("非法请求")
	}
	sql = "select count(id) from goods_shelfs where deleted_at is null and `name`=? and warehouse_id=?  "
	err1 = global.GVA_DB.Raw(sql, info.Name, warehouseId).Scan(&total)
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	if total != 1 {
		return errors.New("仓库不存在该货架")
	}

	err1 = global.GVA_DB.Table("goods_shelfs").Where("name=? and warehouse_id=? and deleted_at is null", info.Name, info.WarehouseId).Update("deleted_at", time.Now())
	if err1.Error != nil {
		return errors.New("系统错误")
	}
	return nil
}
func (warehouseService *WarehouseService) InWarehouse(info request.InWarehouseRequest, sysId uint) (err error) {
	//todo:校验
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		type StaffInfo struct {
			Id          uint
			WarehouseId uint
		}
		var staffInfo StaffInfo
		var inWarehouse warehouse.InWarehouses
		var inWarehouseDetails warehouse.InWarehousesDetails
		sql := "select id,warehouse_id from staffs where deleted_at is null and sys_id=?"
		err1 := tx.Raw(sql, sysId).Scan(&staffInfo)
		if err1.Error != nil {
			return errors.New("系统错误")
		}
		//生成入库单
		now := time.Now()
		orderNumber := "RC" + strconv.FormatInt(now.Unix(), 10)
		inWarehouse = warehouse.InWarehouses{
			OrderNumber: orderNumber,
			WarehouseId: staffInfo.WarehouseId,
			StaffId:     staffInfo.Id,
			Weight:      info.TotalWeight,
			Type:        info.Type,
			FromId:      info.FromId,
			Base: warehouse.Base{
				CreatedAt: now,
			},
		}
		err1 = tx.Table("in_warehouses").Create(&inWarehouse)
		if err1.Error != nil {
			return errors.New("系统错误")
		}
		//生成入库单详情:单号，物品id
		for i := 0; i < len(info.GoodsId); i++ {
			inWarehouseDetails = warehouse.InWarehousesDetails{
				OrderNumber: orderNumber,
				WarehouseId: staffInfo.WarehouseId,
				StaffId:     staffInfo.Id,
				GoodsId:     info.GoodsId[i],
				Weight:      info.Weight[i],
				ShelfName:   info.ShelfName[i],
				Base: warehouse.Base{
					CreatedAt: now,
				},
			}
			err1 = tx.Table("in_warehouses_details").Create(&inWarehouseDetails)
			if err1.Error != nil {
				return errors.New("系统错误")
			}
		}
		//修改货物表货物数量：物品id,warehouse_id
		for i := 0; i < len(info.GoodsId); i++ {
			err1 = tx.Table("goods").Where("id=? and warehouse_id=? and deleted_at is null", info.GoodsId[i], staffInfo.WarehouseId).Update("weight", gorm.Expr("weight+?", info.Weight[i]))
			if err1.Error != nil {
				return errors.New("系统错误")
			}
		}
		//修改货架表货物数量:
		for i := 0; i < len(info.ShelfName); i++ {
			type Weight struct {
				RealTimeWeight int
				MaxWeight      int
			}
			var temWeight Weight
			var goodsId uint
			sql = "select goods_id from goods_shelfs where name=? and warehouse_id=?  and deleted_at is null"
			err1 = tx.Raw(sql, info.ShelfName[i], staffInfo.WarehouseId).Scan(&goodsId)
			if goodsId != info.GoodsId[i] && goodsId != 0 {
				return errors.New("该货架已有货物")
			}
			if err1.Error != nil {
				return errors.New("系统错误")
			}
			sql = "select real_time_weight,max_weight from goods_shelfs where name=? and warehouse_id=? and deleted_at is null"
			err1 = tx.Raw(sql, info.ShelfName[i], staffInfo.WarehouseId).Scan(&temWeight)
			if err1.Error != nil {
				return errors.New("系统错误")
			}
			if temWeight.RealTimeWeight+info.Weight[i] > temWeight.MaxWeight {
				return errors.New("货架货物超出货架最大重量")
			}
			err1 = tx.Table("goods_shelfs").Where("name=? and warehouse_id=? and deleted_at is null", info.ShelfName[i], staffInfo.WarehouseId).Update("real_time_weight", gorm.Expr("real_time_weight+?", info.Weight[i])).Update("goods_id", info.GoodsId[i])
			if err1.Error != nil {
				return errors.New("系统错误")
			}
		}
		return nil
	})

	return err
}

func (warehouseService *WarehouseService) OutWarehouse(info request.OutWarehouseRequest, sysId uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		type StaffInfo struct {
			Id          uint
			WarehouseId uint
		}
		var staffInfo StaffInfo
		var outWarehouse warehouse.OutWarehouses
		var outWarehouseDetails warehouse.OutWarehousesDetails
		sql := "select id,warehouse_id from staffs where deleted_at is null and sys_id=?"
		err1 := tx.Raw(sql, sysId).Scan(&staffInfo)
		if err1.Error != nil {
			return errors.New("系统错误")
		}
		//生成出库单
		now := time.Now()
		orderNumber := "CK" + strconv.FormatInt(now.Unix(), 10)
		outWarehouse = warehouse.OutWarehouses{
			OrderNumber: orderNumber,
			WarehouseId: staffInfo.WarehouseId,
			StaffId:     staffInfo.Id,
			Weight:      info.TotalWeight,
			Type:        info.Type,
			ToId:        info.ToId,
			Base: warehouse.Base{
				CreatedAt: now,
			},
		}
		err1 = tx.Table("out_warehouses").Create(&outWarehouse)
		if err1.Error != nil {
			return errors.New("系统错误")
		}
		//生成出库单详情:单号，物品id
		for i := 0; i < len(info.GoodsId); i++ {
			outWarehouseDetails = warehouse.OutWarehousesDetails{
				OrderNumber: orderNumber,
				WarehouseId: staffInfo.WarehouseId,
				StaffId:     staffInfo.Id,
				GoodsId:     info.GoodsId[i],
				Weight:      info.Weight[i],
				ShelfName:   info.ShelfName[i],
				Base: warehouse.Base{
					CreatedAt: now,
				},
			}
			err1 = tx.Table("out_warehouses_details").Create(&outWarehouseDetails)
			if err1.Error != nil {
				return errors.New("系统错误")
			}
		}
		//修改货物表货物数量：物品id,warehouse_id
		for i := 0; i < len(info.GoodsId); i++ {
			err1 = tx.Table("goods").Where("id=? and warehouse_id=? and deleted_at is null", info.GoodsId[i], staffInfo.WarehouseId).Update("weight", gorm.Expr("weight-?", info.Weight[i]))
			if err1.Error != nil {
				return errors.New("系统错误")
			}
		}
		//修改货架表货物数量:
		for i := 0; i < len(info.ShelfName); i++ {
			type Weight struct {
				Id             uint
				RealTimeWeight int
				MaxWeight      int
			}
			var temWeight Weight
			sql = "select id,real_time_weight,max_weight from goods_shelfs where name=? and warehouse_id=? and goods_id=? and deleted_at is null"
			err1 = tx.Raw(sql, info.ShelfName[i], staffInfo.WarehouseId, info.GoodsId[i]).Scan(&temWeight)
			if err1.Error != nil {
				return errors.New("系统错误")
			}
			if temWeight.Id == 0 {
				goodsName := ""
				sql = "select name from goods where id=? and deleted_at is null"
				err1 = tx.Raw(sql, info.GoodsId[i]).Scan(&goodsName)
				if err1.Error != nil {
					return errors.New("系统错误")
				}
				if goodsName == "" {
					return errors.New("非法请求")
				}
				return errors.New(info.ShelfName[i] + "货架无" + goodsName)
			}
			if temWeight.RealTimeWeight < info.Weight[i] {
				return errors.New(info.ShelfName[i] + "货架货物库存不足，取出失败")
			}
			err1 = tx.Table("goods_shelfs").Where("name=? and warehouse_id=? and deleted_at is null", info.ShelfName[i], staffInfo.WarehouseId).Update("real_time_weight", gorm.Expr("real_time_weight-?", info.Weight[i]))
			if err1.Error != nil {
				return errors.New("系统错误")
			}
		}
		return nil
	})

	return err
}
