package movie

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	MovieRouter
	MovieCategoryRouter
	PlaySourceRouter
	CategoryRouter
}

var (
	mApi             = api.ApiGroupApp.MovieApiGroup.MovieApi
	movieCategoryApi = api.ApiGroupApp.MovieApiGroup.MovieCategoryApi
	playSourceApi    = api.ApiGroupApp.MovieApiGroup.PlaySourceApi
	categoryApi      = api.ApiGroupApp.MovieApiGroup.CategoryApi
)
