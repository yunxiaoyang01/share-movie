package movie

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PlaySourceRouter struct {}

// InitPlaySourceRouter 初始化 playSource表 路由信息
func (s *PlaySourceRouter) InitPlaySourceRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	playSourceRouter := Router.Group("playSource").Use(middleware.OperationRecord())
	playSourceRouterWithoutRecord := Router.Group("playSource")
	playSourceRouterWithoutAuth := PublicRouter.Group("playSource")
	{
		playSourceRouter.POST("createPlaySource", playSourceApi.CreatePlaySource)   // 新建playSource表
		playSourceRouter.DELETE("deletePlaySource", playSourceApi.DeletePlaySource) // 删除playSource表
		playSourceRouter.DELETE("deletePlaySourceByIds", playSourceApi.DeletePlaySourceByIds) // 批量删除playSource表
		playSourceRouter.PUT("updatePlaySource", playSourceApi.UpdatePlaySource)    // 更新playSource表
	}
	{
		playSourceRouterWithoutRecord.GET("findPlaySource", playSourceApi.FindPlaySource)        // 根据ID获取playSource表
		playSourceRouterWithoutRecord.GET("getPlaySourceList", playSourceApi.GetPlaySourceList)  // 获取playSource表列表
	}
	{
	    playSourceRouterWithoutAuth.GET("getPlaySourcePublic", playSourceApi.GetPlaySourcePublic)  // playSource表开放接口
	}
}
