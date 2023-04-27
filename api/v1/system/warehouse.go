package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/warehouse/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type WarehouseApi struct{}

// todo:分页、条件筛选
func (w *WarehouseApi) GetV1WarehousesList(c *gin.Context) {
	var pageInfo request.Page
	pageInfo.Page, _ = strconv.Atoi(c.Query("page"))
	pageInfo.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	date, total, err := warehouseService.GetV1WarehousesList(pageInfo)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     date,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "成功", c)
	}
}
func (w *WarehouseApi) GetV1GoodsList(c *gin.Context) {
	var pageInfo request.Page
	pageInfo.Page, _ = strconv.Atoi(c.Query("page"))
	pageInfo.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	date, total, err := warehouseService.GetV1GoodsList(pageInfo)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     date,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "成功", c)
	}
}
func (w *WarehouseApi) GetV1StaffsList(c *gin.Context) {
	var pageInfo request.Page
	pageInfo.Page, _ = strconv.Atoi(c.Query("page"))
	pageInfo.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	date, total, err := warehouseService.GetV1StaffsList(pageInfo)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     date,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "成功", c)
	}
}
func (w *WarehouseApi) GetV1CustomersList(c *gin.Context) {
	var pageInfo request.Page
	pageInfo.Page, _ = strconv.Atoi(c.Query("page"))
	pageInfo.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	date, total, err := warehouseService.GetV1CustomersList(pageInfo)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     date,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "成功", c)
	}
}
func (w *WarehouseApi) GetV1SuppliersList(c *gin.Context) {
	var pageInfo request.Page
	pageInfo.Page, _ = strconv.Atoi(c.Query("page"))
	pageInfo.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	date, total, err := warehouseService.GetV1SuppliersList(pageInfo)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     date,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "成功", c)
	}
}
func (w *WarehouseApi) GetV1GoodsShelfsList(c *gin.Context) {
	var pageInfo request.Page
	pageInfo.Page, _ = strconv.Atoi(c.Query("page"))
	pageInfo.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	date, total, err := warehouseService.GetV1GoodsShelfsList(pageInfo)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     date,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "成功", c)
	}
}
func (w *WarehouseApi) GetV1OutWarehousesList(c *gin.Context) {
	var pageInfo request.Page
	pageInfo.Page, _ = strconv.Atoi(c.Query("page"))
	pageInfo.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	date, total, err := warehouseService.GetV1OutWarehousesList(pageInfo)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     date,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "成功", c)
	}
}
func (w *WarehouseApi) GetV1OutWarehousesDetail(c *gin.Context) {
	var pageInfo request.Page
	pageInfo.Page, _ = strconv.Atoi(c.Query("page"))
	pageInfo.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	orderNumber := c.Query("orderNumber")
	date, total, err := warehouseService.GetV1OutWarehousesDetail(orderNumber, pageInfo)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     date,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "成功", c)
	}
}
func (w *WarehouseApi) GetV1InWarehousesList(c *gin.Context) {
	var pageInfo request.Page
	pageInfo.Page, _ = strconv.Atoi(c.Query("page"))
	pageInfo.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	date, total, err := warehouseService.GetV1InWarehousesList(pageInfo)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     date,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "成功", c)
	}
}
func (w *WarehouseApi) GetV1InWarehousesDetail(c *gin.Context) {
	var pageInfo request.Page
	pageInfo.Page, _ = strconv.Atoi(c.Query("page"))
	pageInfo.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	orderNumber := c.Query("orderNumber")
	date, total, err := warehouseService.GetV1InWarehousesDetail(orderNumber, pageInfo)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     date,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "成功", c)
	}
}
func (w *WarehouseApi) AddCustomer(c *gin.Context) {
	var info request.CustomersRequest
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithDetailed("", "绑定失败", c)
		return
	}
	if info.Phone == "" && info.Email == "" {
		response.FailWithDetailed("", "必须要有邮箱或手机", c)
		return
	}
	err = warehouseService.AddCustomer(info)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed("", "成功", c)
	}
}
func (w *WarehouseApi) AddSupplier(c *gin.Context) {
	var info request.SuppliersRequest
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithDetailed("", "绑定失败", c)
		return
	}
	if info.Phone == "" && info.Email == "" {
		response.FailWithDetailed("", "必须要有邮箱或手机", c)
		return
	}
	err = warehouseService.AddSupplier(info)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed("", "成功", c)
	}
}
func (w *WarehouseApi) AddWarehouse(c *gin.Context) {
	var info request.WarehousesRequest
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithDetailed("", "绑定失败", c)
		return
	}
	err = warehouseService.AddWarehouse(info)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed("", "成功", c)
	}
}
func (w *WarehouseApi) AddWarehouseAndstaffRelations(c *gin.Context) {
	var info request.AddWarehouseAndstaffRelations
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithDetailed("", "绑定失败", c)
		return
	}
	err = warehouseService.AddWarehouseAndstaffRelations(info)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed("", "成功", c)
	}
}
func (w *WarehouseApi) CancelWarehouseAndstaffRelations(c *gin.Context) {
	var info request.CancelWarehouseAndstaffRelations
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithDetailed("", "绑定失败", c)
		return
	}
	err = warehouseService.CancelWarehouseAndstaffRelations(info)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed("", "成功", c)
	}
}
func (w *WarehouseApi) AddStaff(c *gin.Context) {
	var r request.AddStaffRequest
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(r, utils.Register1Verify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var authorities []system.SysAuthority
	authorities = append(authorities, system.SysAuthority{
		AuthorityId: 2,
	})
	user := &system.SysUser{Username: r.Username, NickName: r.NickName, Password: r.Password, HeaderImg: r.HeaderImg, AuthorityId: 2, Authorities: authorities, Enable: r.Enable, Phone: r.Phone, Email: r.Email}
	userReturn, err := warehouseService.AddStaff(*user, r.WarehouseId)
	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)
}
func (w *WarehouseApi) DeleteStaff(c *gin.Context) {
	var info request.DeleteStaff
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithDetailed("", "绑定失败", c)
		return
	}
	err = warehouseService.DeleteStaff(info)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed("", "成功", c)
	}
}
func (w *WarehouseApi) DeleteCustomer(c *gin.Context) {
	var info request.DeleteCustomer
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithDetailed("", "绑定失败", c)
		return
	}
	err = warehouseService.DeleteCustomer(info)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed("", "成功", c)
	}
}
func (w *WarehouseApi) DeleteSupplier(c *gin.Context) {
	var info request.DeleteSupplier
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithDetailed("", "绑定失败", c)
		return
	}
	err = warehouseService.DeleteSupplier(info)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed("", "成功", c)
	}
}
func (w *WarehouseApi) DeleteWarehouse(c *gin.Context) {
	var info request.DeleteWarehouse
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithDetailed("", "绑定失败", c)
		return
	}
	err = warehouseService.DeleteWarehouse(info)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed("", "成功", c)
	}
}

func (w *WarehouseApi) UpdateStaff(c *gin.Context) {
	var info request.UpdateStaffRequest
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithDetailed("", "绑定失败", c)
		return
	}
	if info.Phone == "" && info.Email == "" {
		response.FailWithDetailed("", "必须要有邮箱或手机", c)
		return
	}
	err = warehouseService.UpdateStaff(info)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed("", "成功", c)
	}
}
func (w *WarehouseApi) UpdateCustomer(c *gin.Context) {
	var info request.UpdateCustomerRequest
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithDetailed("", "绑定失败", c)
		return
	}
	if info.Phone == "" && info.Email == "" {
		response.FailWithDetailed("", "必须要有邮箱或手机", c)
		return
	}
	err = warehouseService.UpdateCustomer(info)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed("", "成功", c)
	}
}
func (w *WarehouseApi) UpdateSupplier(c *gin.Context) {
	var info request.UpdateSupplierRequest
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithDetailed("", "绑定失败", c)
		return
	}
	if info.Phone == "" && info.Email == "" {
		response.FailWithDetailed("", "必须要有邮箱或手机", c)
		return
	}
	err = warehouseService.UpdateSupplier(info)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed("", "成功", c)
	}
}
func (w *WarehouseApi) UpdateWarehouse(c *gin.Context) {
	var info request.UpdateWarehouseRequest
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithDetailed("", "绑定失败", c)
		return
	}
	err = warehouseService.UpdateWarehouse(info)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed("", "成功", c)
	}
}

// -------------------------v2-----------------
func (w *WarehouseApi) GetV2GoodsList(c *gin.Context) {
	var pageInfo request.Page
	pageInfo.Page, _ = strconv.Atoi(c.Query("page"))
	pageInfo.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	date, total, err := warehouseService.GetV2GoodsList(utils.GetUserID(c), pageInfo)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     date,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "成功", c)
	}
}
func (w *WarehouseApi) GetV2StaffsList(c *gin.Context) {
	var pageInfo request.Page
	pageInfo.Page, _ = strconv.Atoi(c.Query("page"))
	pageInfo.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	date, total, err := warehouseService.GetV2StaffsList(utils.GetUserID(c), pageInfo)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     date,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "成功", c)
	}
}
func (w *WarehouseApi) GetV2CustomersList(c *gin.Context) {
	var pageInfo request.Page
	pageInfo.Page, _ = strconv.Atoi(c.Query("page"))
	pageInfo.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	date, total, err := warehouseService.GetV2CustomersList(pageInfo)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     date,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "成功", c)
	}
}
func (w *WarehouseApi) GetV2SuppliersList(c *gin.Context) {
	var pageInfo request.Page
	pageInfo.Page, _ = strconv.Atoi(c.Query("page"))
	pageInfo.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	date, total, err := warehouseService.GetV2SuppliersList(pageInfo)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     date,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "成功", c)
	}
}
func (w *WarehouseApi) GetV2GoodsShelfsList(c *gin.Context) {
	var pageInfo request.Page
	pageInfo.Page, _ = strconv.Atoi(c.Query("page"))
	pageInfo.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	date, total, err := warehouseService.GetV2GoodsShelfsList(utils.GetUserID(c), pageInfo)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     date,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "成功", c)
	}
}
func (w *WarehouseApi) GetV2OutWarehousesList(c *gin.Context) {
	var pageInfo request.Page
	pageInfo.Page, _ = strconv.Atoi(c.Query("page"))
	pageInfo.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	date, total, err := warehouseService.GetV2OutWarehousesList(utils.GetUserID(c), pageInfo)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     date,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "成功", c)
	}
}
func (w *WarehouseApi) GetV2OutWarehousesDetail(c *gin.Context) {
	orderNumber := c.Query("orderNumber")
	var pageInfo request.Page
	pageInfo.Page, _ = strconv.Atoi(c.Query("page"))
	pageInfo.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	date, total, err := warehouseService.GetV2OutWarehousesDetail(orderNumber, utils.GetUserID(c), pageInfo)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     date,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "成功", c)
	}
}
func (w *WarehouseApi) GetV2InWarehousesList(c *gin.Context) {
	var pageInfo request.Page
	pageInfo.Page, _ = strconv.Atoi(c.Query("page"))
	pageInfo.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	date, total, err := warehouseService.GetV2InWarehousesList(utils.GetUserID(c), pageInfo)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     date,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "成功", c)
	}
}
func (w *WarehouseApi) GetV2InWarehousesDetail(c *gin.Context) {
	orderNumber := c.Query("orderNumber")
	var pageInfo request.Page
	pageInfo.Page, _ = strconv.Atoi(c.Query("page"))
	pageInfo.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	date, total, err := warehouseService.GetV2InWarehousesDetail(orderNumber, utils.GetUserID(c), pageInfo)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     date,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "成功", c)
	}
}

func (w *WarehouseApi) GetV2WarehousesList(c *gin.Context) {
	var pageInfo request.Page
	pageInfo.Page, _ = strconv.Atoi(c.Query("page"))
	pageInfo.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	date, total, err := warehouseService.GetV2WarehousesList(pageInfo)
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     date,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "成功", c)
	}
}
func (w *WarehouseApi) AddGood(c *gin.Context) {
	var info request.AddGoodsRequest
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithDetailed("", "绑定失败", c)
		return
	}
	err = warehouseService.AddGood(info, utils.GetUserID(c))
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed("", "成功", c)
	}
}
func (w *WarehouseApi) AddGoodsShelf(c *gin.Context) {
	var info request.AddGoodsShelfRequest
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithDetailed("", "绑定失败", c)
		return
	}
	err = warehouseService.AddGoodsShelf(info, utils.GetUserID(c))
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed("", "成功", c)
	}
}
func (w *WarehouseApi) DeleteGood(c *gin.Context) {
	var info request.DeleteGoodRequest
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithDetailed("", "绑定失败", c)
		return
	}
	err = warehouseService.DeleteGood(info, utils.GetUserID(c))
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed("", "成功", c)
	}
}
func (w *WarehouseApi) DeleteGoodsShelf(c *gin.Context) {
	var info request.DeleteGoodsShelfRequest
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithDetailed("", "绑定失败", c)
		return
	}
	err = warehouseService.DeleteGoodsShelf(info, utils.GetUserID(c))
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed("", "成功", c)
	}
}
func (w *WarehouseApi) InWarehouse(c *gin.Context) {
	var info request.InWarehouseRequest
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithDetailed("", "绑定失败", c)
		return
	}
	if info.FromId == 0 {
		response.FailWithDetailed("", "来源不能为空", c)
		return
	}
	if info.TotalWeight <= 0 {
		response.FailWithDetailed("", "入库数量不能为零", c)
		return
	}
	err = warehouseService.InWarehouse(info, utils.GetUserID(c))
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed("", "成功", c)
	}
}
func (w *WarehouseApi) OutWarehouse(c *gin.Context) {
	var info request.OutWarehouseRequest
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithDetailed("", "绑定失败", c)
		return
	}
	if info.ToId == 0 {
		response.FailWithDetailed("", "去处不能为空", c)
		return
	}
	if info.TotalWeight <= 0 {
		response.FailWithDetailed("", "出库数量不能为零", c)
		return
	}
	err = warehouseService.OutWarehouse(info, utils.GetUserID(c))
	if err != nil {
		response.FailWithDetailed("", err.Error(), c)
	} else {
		response.OkWithDetailed("", "成功", c)
	}
}
