
package movie

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/movie"
    movieReq "github.com/flipped-aurora/gin-vue-admin/server/model/movie/request"
)

type MovieCategoryService struct {}
// CreateMovieCategory 创建movieCategory表记录
// Author [yourname](https://github.com/yourname)
func (movieCategoryService *MovieCategoryService) CreateMovieCategory(movieCategory *movie.MovieCategory) (err error) {
	err = global.GVA_DB.Create(movieCategory).Error
	return err
}

// DeleteMovieCategory 删除movieCategory表记录
// Author [yourname](https://github.com/yourname)
func (movieCategoryService *MovieCategoryService)DeleteMovieCategory(movieId string) (err error) {
	err = global.GVA_DB.Delete(&movie.MovieCategory{},"movie_id = ?",movieId).Error
	return err
}

// DeleteMovieCategoryByIds 批量删除movieCategory表记录
// Author [yourname](https://github.com/yourname)
func (movieCategoryService *MovieCategoryService)DeleteMovieCategoryByIds(movieIds []string) (err error) {
	err = global.GVA_DB.Delete(&[]movie.MovieCategory{},"movie_id in ?",movieIds).Error
	return err
}

// UpdateMovieCategory 更新movieCategory表记录
// Author [yourname](https://github.com/yourname)
func (movieCategoryService *MovieCategoryService)UpdateMovieCategory(movieCategory movie.MovieCategory) (err error) {
	err = global.GVA_DB.Model(&movie.MovieCategory{}).Where("movie_id = ?",movieCategory.MovieId).Updates(&movieCategory).Error
	return err
}

// GetMovieCategory 根据movieId获取movieCategory表记录
// Author [yourname](https://github.com/yourname)
func (movieCategoryService *MovieCategoryService)GetMovieCategory(movieId string) (movieCategory movie.MovieCategory, err error) {
	err = global.GVA_DB.Where("movie_id = ?", movieId).First(&movieCategory).Error
	return
}
// GetMovieCategoryInfoList 分页获取movieCategory表记录
// Author [yourname](https://github.com/yourname)
func (movieCategoryService *MovieCategoryService)GetMovieCategoryInfoList(info movieReq.MovieCategorySearch) (list []movie.MovieCategory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&movie.MovieCategory{})
    var movieCategorys []movie.MovieCategory
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&movieCategorys).Error
	return  movieCategorys, total, err
}
func (movieCategoryService *MovieCategoryService)GetMovieCategoryPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
