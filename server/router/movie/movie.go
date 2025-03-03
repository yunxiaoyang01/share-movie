package movie

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MovieRouter struct {}

// InitMovieRouter 初始化 movie表 路由信息
func (s *MovieRouter) InitMovieRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	mRouter := Router.Group("m").Use(middleware.OperationRecord())
	mRouterWithoutRecord := Router.Group("m")
	mRouterWithoutAuth := PublicRouter.Group("m")
	{
		mRouter.POST("createMovie", mApi.CreateMovie)   // 新建movie表
		mRouter.DELETE("deleteMovie", mApi.DeleteMovie) // 删除movie表
		mRouter.DELETE("deleteMovieByIds", mApi.DeleteMovieByIds) // 批量删除movie表
		mRouter.PUT("updateMovie", mApi.UpdateMovie)    // 更新movie表
	}
	{
		mRouterWithoutRecord.GET("findMovie", mApi.FindMovie)        // 根据ID获取movie表
		mRouterWithoutRecord.GET("getMovieList", mApi.GetMovieList)  // 获取movie表列表
	}
	{
	    mRouterWithoutAuth.GET("getMoviePublic", mApi.GetMoviePublic)  // movie表开放接口
	}
}
