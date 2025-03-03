package movie

type ServiceGroup struct {
	MovieService
	MovieCategoryService
	PlaySourceService
	CategoryService
}
