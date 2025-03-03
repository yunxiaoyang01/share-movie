package movie

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MovieCategoryRouter struct {}

// InitMovieCategoryRouter 初始化 movieCategory表 路由信息
func (s *MovieCategoryRouter) InitMovieCategoryRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	movieCategoryRouter := Router.Group("movieCategory").Use(middleware.OperationRecord())
	movieCategoryRouterWithoutRecord := Router.Group("movieCategory")
	movieCategoryRouterWithoutAuth := PublicRouter.Group("movieCategory")
	{
		movieCategoryRouter.POST("createMovieCategory", movieCategoryApi.CreateMovieCategory)   // 新建movieCategory表
		movieCategoryRouter.DELETE("deleteMovieCategory", movieCategoryApi.DeleteMovieCategory) // 删除movieCategory表
		movieCategoryRouter.DELETE("deleteMovieCategoryByIds", movieCategoryApi.DeleteMovieCategoryByIds) // 批量删除movieCategory表
		movieCategoryRouter.PUT("updateMovieCategory", movieCategoryApi.UpdateMovieCategory)    // 更新movieCategory表
	}
	{
		movieCategoryRouterWithoutRecord.GET("findMovieCategory", movieCategoryApi.FindMovieCategory)        // 根据ID获取movieCategory表
		movieCategoryRouterWithoutRecord.GET("getMovieCategoryList", movieCategoryApi.GetMovieCategoryList)  // 获取movieCategory表列表
	}
	{
	    movieCategoryRouterWithoutAuth.GET("getMovieCategoryPublic", movieCategoryApi.GetMovieCategoryPublic)  // movieCategory表开放接口
	}
}
