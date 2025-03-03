import service from '@/utils/request'
// @Tags Movie
// @Summary 创建movie表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Movie true "创建movie表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /m/createMovie [post]
export const createMovie = (data) => {
  return service({
    url: '/m/createMovie',
    method: 'post',
    data
  })
}

// @Tags Movie
// @Summary 删除movie表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Movie true "删除movie表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /m/deleteMovie [delete]
export const deleteMovie = (params) => {
  return service({
    url: '/m/deleteMovie',
    method: 'delete',
    params
  })
}

// @Tags Movie
// @Summary 批量删除movie表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除movie表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /m/deleteMovie [delete]
export const deleteMovieByIds = (params) => {
  return service({
    url: '/m/deleteMovieByIds',
    method: 'delete',
    params
  })
}

// @Tags Movie
// @Summary 更新movie表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Movie true "更新movie表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /m/updateMovie [put]
export const updateMovie = (data) => {
  return service({
    url: '/m/updateMovie',
    method: 'put',
    data
  })
}

// @Tags Movie
// @Summary 用id查询movie表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Movie true "用id查询movie表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /m/findMovie [get]
export const findMovie = (params) => {
  return service({
    url: '/m/findMovie',
    method: 'get',
    params
  })
}

// @Tags Movie
// @Summary 分页获取movie表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取movie表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /m/getMovieList [get]
export const getMovieList = (params) => {
  return service({
    url: '/m/getMovieList',
    method: 'get',
    params
  })
}

// @Tags Movie
// @Summary 不需要鉴权的movie表接口
// @Accept application/json
// @Produce application/json
// @Param data query movieReq.MovieSearch true "分页获取movie表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /m/getMoviePublic [get]
export const getMoviePublic = () => {
  return service({
    url: '/m/getMoviePublic',
    method: 'get',
  })
}
