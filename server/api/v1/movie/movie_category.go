package movie

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/movie"
    movieReq "github.com/flipped-aurora/gin-vue-admin/server/model/movie/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type MovieCategoryApi struct {}



// CreateMovieCategory 创建movieCategory表
// @Tags MovieCategory
// @Summary 创建movieCategory表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body movie.MovieCategory true "创建movieCategory表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /movieCategory/createMovieCategory [post]
func (movieCategoryApi *MovieCategoryApi) CreateMovieCategory(c *gin.Context) {
	var movieCategory movie.MovieCategory
	err := c.ShouldBindJSON(&movieCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = movieCategoryService.CreateMovieCategory(&movieCategory)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteMovieCategory 删除movieCategory表
// @Tags MovieCategory
// @Summary 删除movieCategory表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body movie.MovieCategory true "删除movieCategory表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /movieCategory/deleteMovieCategory [delete]
func (movieCategoryApi *MovieCategoryApi) DeleteMovieCategory(c *gin.Context) {
	movieId := c.Query("movieId")
	err := movieCategoryService.DeleteMovieCategory(movieId)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteMovieCategoryByIds 批量删除movieCategory表
// @Tags MovieCategory
// @Summary 批量删除movieCategory表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /movieCategory/deleteMovieCategoryByIds [delete]
func (movieCategoryApi *MovieCategoryApi) DeleteMovieCategoryByIds(c *gin.Context) {
	movieIds := c.QueryArray("movieIds[]")
	err := movieCategoryService.DeleteMovieCategoryByIds(movieIds)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateMovieCategory 更新movieCategory表
// @Tags MovieCategory
// @Summary 更新movieCategory表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body movie.MovieCategory true "更新movieCategory表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /movieCategory/updateMovieCategory [put]
func (movieCategoryApi *MovieCategoryApi) UpdateMovieCategory(c *gin.Context) {
	var movieCategory movie.MovieCategory
	err := c.ShouldBindJSON(&movieCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = movieCategoryService.UpdateMovieCategory(movieCategory)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindMovieCategory 用id查询movieCategory表
// @Tags MovieCategory
// @Summary 用id查询movieCategory表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param movieId query int true "用id查询movieCategory表"
// @Success 200 {object} response.Response{data=movie.MovieCategory,msg=string} "查询成功"
// @Router /movieCategory/findMovieCategory [get]
func (movieCategoryApi *MovieCategoryApi) FindMovieCategory(c *gin.Context) {
	movieId := c.Query("movieId")
	removieCategory, err := movieCategoryService.GetMovieCategory(movieId)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(removieCategory, c)
}
// GetMovieCategoryList 分页获取movieCategory表列表
// @Tags MovieCategory
// @Summary 分页获取movieCategory表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query movieReq.MovieCategorySearch true "分页获取movieCategory表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /movieCategory/getMovieCategoryList [get]
func (movieCategoryApi *MovieCategoryApi) GetMovieCategoryList(c *gin.Context) {
	var pageInfo movieReq.MovieCategorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := movieCategoryService.GetMovieCategoryInfoList(pageInfo)
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

// GetMovieCategoryPublic 不需要鉴权的movieCategory表接口
// @Tags MovieCategory
// @Summary 不需要鉴权的movieCategory表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /movieCategory/getMovieCategoryPublic [get]
func (movieCategoryApi *MovieCategoryApi) GetMovieCategoryPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    movieCategoryService.GetMovieCategoryPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的movieCategory表接口信息",
    }, "获取成功", c)
}
