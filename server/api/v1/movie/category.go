package movie

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/movie"
    movieReq "github.com/flipped-aurora/gin-vue-admin/server/model/movie/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type CategoryApi struct {}



// CreateCategory 创建category表
// @Tags Category
// @Summary 创建category表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body movie.Category true "创建category表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /category/createCategory [post]
func (categoryApi *CategoryApi) CreateCategory(c *gin.Context) {
	var category movie.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = categoryService.CreateCategory(&category)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteCategory 删除category表
// @Tags Category
// @Summary 删除category表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body movie.Category true "删除category表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /category/deleteCategory [delete]
func (categoryApi *CategoryApi) DeleteCategory(c *gin.Context) {
	id := c.Query("id")
	err := categoryService.DeleteCategory(id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCategoryByIds 批量删除category表
// @Tags Category
// @Summary 批量删除category表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /category/deleteCategoryByIds [delete]
func (categoryApi *CategoryApi) DeleteCategoryByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := categoryService.DeleteCategoryByIds(ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateCategory 更新category表
// @Tags Category
// @Summary 更新category表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body movie.Category true "更新category表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /category/updateCategory [put]
func (categoryApi *CategoryApi) UpdateCategory(c *gin.Context) {
	var category movie.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = categoryService.UpdateCategory(category)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindCategory 用id查询category表
// @Tags Category
// @Summary 用id查询category表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询category表"
// @Success 200 {object} response.Response{data=movie.Category,msg=string} "查询成功"
// @Router /category/findCategory [get]
func (categoryApi *CategoryApi) FindCategory(c *gin.Context) {
	id := c.Query("id")
	recategory, err := categoryService.GetCategory(id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(recategory, c)
}
// GetCategoryList 分页获取category表列表
// @Tags Category
// @Summary 分页获取category表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query movieReq.CategorySearch true "分页获取category表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /category/getCategoryList [get]
func (categoryApi *CategoryApi) GetCategoryList(c *gin.Context) {
	var pageInfo movieReq.CategorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := categoryService.GetCategoryInfoList(pageInfo)
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

// GetCategoryPublic 不需要鉴权的category表接口
// @Tags Category
// @Summary 不需要鉴权的category表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /category/getCategoryPublic [get]
func (categoryApi *CategoryApi) GetCategoryPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    categoryService.GetCategoryPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的category表接口信息",
    }, "获取成功", c)
}
