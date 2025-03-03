package movie

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CategoryRouter struct {}

// InitCategoryRouter 初始化 category表 路由信息
func (s *CategoryRouter) InitCategoryRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	categoryRouter := Router.Group("category").Use(middleware.OperationRecord())
	categoryRouterWithoutRecord := Router.Group("category")
	categoryRouterWithoutAuth := PublicRouter.Group("category")
	{
		categoryRouter.POST("createCategory", categoryApi.CreateCategory)   // 新建category表
		categoryRouter.DELETE("deleteCategory", categoryApi.DeleteCategory) // 删除category表
		categoryRouter.DELETE("deleteCategoryByIds", categoryApi.DeleteCategoryByIds) // 批量删除category表
		categoryRouter.PUT("updateCategory", categoryApi.UpdateCategory)    // 更新category表
	}
	{
		categoryRouterWithoutRecord.GET("findCategory", categoryApi.FindCategory)        // 根据ID获取category表
		categoryRouterWithoutRecord.GET("getCategoryList", categoryApi.GetCategoryList)  // 获取category表列表
	}
	{
	    categoryRouterWithoutAuth.GET("getCategoryPublic", categoryApi.GetCategoryPublic)  // category表开放接口
	}
}
