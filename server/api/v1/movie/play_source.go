package movie

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/movie"
    movieReq "github.com/flipped-aurora/gin-vue-admin/server/model/movie/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type PlaySourceApi struct {}



// CreatePlaySource 创建playSource表
// @Tags PlaySource
// @Summary 创建playSource表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body movie.PlaySource true "创建playSource表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /playSource/createPlaySource [post]
func (playSourceApi *PlaySourceApi) CreatePlaySource(c *gin.Context) {
	var playSource movie.PlaySource
	err := c.ShouldBindJSON(&playSource)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = playSourceService.CreatePlaySource(&playSource)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeletePlaySource 删除playSource表
// @Tags PlaySource
// @Summary 删除playSource表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body movie.PlaySource true "删除playSource表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /playSource/deletePlaySource [delete]
func (playSourceApi *PlaySourceApi) DeletePlaySource(c *gin.Context) {
	id := c.Query("id")
	err := playSourceService.DeletePlaySource(id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePlaySourceByIds 批量删除playSource表
// @Tags PlaySource
// @Summary 批量删除playSource表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /playSource/deletePlaySourceByIds [delete]
func (playSourceApi *PlaySourceApi) DeletePlaySourceByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := playSourceService.DeletePlaySourceByIds(ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePlaySource 更新playSource表
// @Tags PlaySource
// @Summary 更新playSource表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body movie.PlaySource true "更新playSource表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /playSource/updatePlaySource [put]
func (playSourceApi *PlaySourceApi) UpdatePlaySource(c *gin.Context) {
	var playSource movie.PlaySource
	err := c.ShouldBindJSON(&playSource)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = playSourceService.UpdatePlaySource(playSource)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindPlaySource 用id查询playSource表
// @Tags PlaySource
// @Summary 用id查询playSource表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询playSource表"
// @Success 200 {object} response.Response{data=movie.PlaySource,msg=string} "查询成功"
// @Router /playSource/findPlaySource [get]
func (playSourceApi *PlaySourceApi) FindPlaySource(c *gin.Context) {
	id := c.Query("id")
	replaySource, err := playSourceService.GetPlaySource(id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(replaySource, c)
}
// GetPlaySourceList 分页获取playSource表列表
// @Tags PlaySource
// @Summary 分页获取playSource表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query movieReq.PlaySourceSearch true "分页获取playSource表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /playSource/getPlaySourceList [get]
func (playSourceApi *PlaySourceApi) GetPlaySourceList(c *gin.Context) {
	var pageInfo movieReq.PlaySourceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := playSourceService.GetPlaySourceInfoList(pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetPlaySourcePublic 不需要鉴权的playSource表接口
// @Tags PlaySource
// @Summary 不需要鉴权的playSource表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /playSource/getPlaySourcePublic [get]
func (playSourceApi *PlaySourceApi) GetPlaySourcePublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    playSourceService.GetPlaySourcePublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的playSource表接口信息",
    }, "获取成功", c)
}
