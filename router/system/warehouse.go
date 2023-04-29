package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type WarehouseRouter struct{}

func (s *WarehouseRouter) InitWarehouseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	warehouseRouter := Router.Group("warehouse")
	warehouseApi := v1.ApiGroupApp.SystemApiGroup.WarehouseApi
	{
		//一级用户
		warehouseRouter.GET("v1/getWarehousesList", warehouseApi.GetV1WarehousesList)           //获取仓库列表
		warehouseRouter.GET("v1/getGoodsList", warehouseApi.GetV1GoodsList)                     //获取商品列表
		warehouseRouter.GET("v1/getStaffsList", warehouseApi.GetV1StaffsList)                   //获取员工列表
		warehouseRouter.GET("v1/getCustomersList", warehouseApi.GetV1CustomersList)             //获取客户列表
		warehouseRouter.GET("v1/getSuppliersList", warehouseApi.GetV1SuppliersList)             //获取供应商列表
		warehouseRouter.GET("v1/getGoodsShelfsList", warehouseApi.GetV1GoodsShelfsList)         //获取货架列表
		warehouseRouter.GET("v1/getOutWarehousesList", warehouseApi.GetV1OutWarehousesList)     //获取出库单列表
		warehouseRouter.GET("v1/getOutWarehousesDetail", warehouseApi.GetV1OutWarehousesDetail) //获取出库单详情
		warehouseRouter.GET("v1/getInWarehousesList", warehouseApi.GetV1InWarehousesList)       //获取入库单列表
		warehouseRouter.GET("v1/getInWarehousesDetail", warehouseApi.GetV1InWarehousesDetail)   //获取入库单详情

		warehouseRouter.POST("v1/addCustomer", warehouseApi.AddCustomer)                          //增加客户
		warehouseRouter.POST("v1/addSupplier", warehouseApi.AddSupplier)                          //增加供应商
		warehouseRouter.POST("v1/addWarehouse", warehouseApi.AddWarehouse)                        //增加仓库
		warehouseRouter.POST("v1/addRelations", warehouseApi.AddWarehouseAndstaffRelations)       //增加员工与仓库的关系
		warehouseRouter.POST("v1/cancelRelations", warehouseApi.CancelWarehouseAndstaffRelations) //解除员工与仓库的关系
		warehouseRouter.POST("v1/addStaff", warehouseApi.AddStaff)                                //增加员工
		warehouseRouter.POST("v1/deleteStaff", warehouseApi.DeleteStaff)                          //删除员工
		warehouseRouter.POST("v1/deleteCustomer", warehouseApi.DeleteCustomer)                    //删除客户
		warehouseRouter.POST("v1/deleteSupplier", warehouseApi.DeleteSupplier)                    //删除供应商
		warehouseRouter.POST("v1/deleteWarehouse", warehouseApi.DeleteWarehouse)                  //删除仓库

		warehouseRouter.POST("v1/updateStaff", warehouseApi.UpdateStaff)         //修改员工信息
		warehouseRouter.POST("v1/updateCustomer", warehouseApi.UpdateCustomer)   //修改客户信息
		warehouseRouter.POST("v1/updateSupplier", warehouseApi.UpdateSupplier)   //修改供应商信息
		warehouseRouter.POST("v1/updateWarehouse", warehouseApi.UpdateWarehouse) //修改仓库信息
		warehouseRouter.POST("v1/resetPassword", warehouseApi.ResetPassword)     //重置员工密码
		//todo:修改出库单,修改出库单详情,修改入库单,修改入库单详情

		//二级用户
		warehouseRouter.GET("v2/getGoodsList", warehouseApi.GetV2GoodsList)                     //获取商品列表
		warehouseRouter.GET("v2/getStaffsList", warehouseApi.GetV2StaffsList)                   //获取员工列表
		warehouseRouter.GET("v2/getCustomersList", warehouseApi.GetV2CustomersList)             //获取客户列表
		warehouseRouter.GET("v2/getSuppliersList", warehouseApi.GetV2SuppliersList)             //获取供应商列表
		warehouseRouter.GET("v2/getGoodsShelfsList", warehouseApi.GetV2GoodsShelfsList)         //获取货架列表
		warehouseRouter.GET("v2/getOutWarehousesList", warehouseApi.GetV2OutWarehousesList)     //获取出库单列表
		warehouseRouter.GET("v2/getOutWarehousesDetail", warehouseApi.GetV2OutWarehousesDetail) //获取出库单详情
		warehouseRouter.GET("v2/getInWarehousesList", warehouseApi.GetV2InWarehousesList)       //获取入库单列表
		warehouseRouter.GET("v2/getInWarehousesDetail", warehouseApi.GetV2InWarehousesDetail)   //获取入库单详情
		warehouseRouter.GET("v2/getWarehousesList", warehouseApi.GetV2WarehousesList)           //生成出/入库列表-获取仓库列表

		warehouseRouter.POST("v2/addGood", warehouseApi.AddGood)                   //增加商品
		warehouseRouter.POST("v2/addGoodsShelf", warehouseApi.AddGoodsShelf)       //增加货架
		warehouseRouter.POST("v2/deleteGood", warehouseApi.DeleteGood)             //删除商品
		warehouseRouter.POST("v2/deleteGoodsShelf", warehouseApi.DeleteGoodsShelf) //删除货架
		warehouseRouter.POST("v2/inWarehouse", warehouseApi.InWarehouse)           //生成入库单与入库单详情
		warehouseRouter.POST("v2/outWarehouse", warehouseApi.OutWarehouse)         //生成出库单与出库单详情
	}
	return warehouseRouter
}
