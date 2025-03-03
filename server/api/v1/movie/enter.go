package movie

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	MovieApi
	MovieCategoryApi
	PlaySourceApi
	CategoryApi
}

var (
	mService             = service.ServiceGroupApp.MovieServiceGroup.MovieService
	movieCategoryService = service.ServiceGroupApp.MovieServiceGroup.MovieCategoryService
	playSourceService    = service.ServiceGroupApp.MovieServiceGroup.PlaySourceService
	categoryService      = service.ServiceGroupApp.MovieServiceGroup.CategoryService
)
