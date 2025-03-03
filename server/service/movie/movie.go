package movie

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/movie"
	movieReq "github.com/flipped-aurora/gin-vue-admin/server/model/movie/request"
)

type MovieService struct{}

// CreateMovie 创建movie表记录
// Author [yourname](https://github.com/yourname)
func (mService *MovieService) CreateMovie(m *movie.Movie) (err error) {
	err = global.GVA_DB.Create(m).Error
	return err
}

// DeleteMovie 删除movie表记录
// Author [yourname](https://github.com/yourname)
func (mService *MovieService) DeleteMovie(id string) (err error) {
	err = global.GVA_DB.Delete(&movie.Movie{}, "id = ?", id).Error
	return err
}

// DeleteMovieByIds 批量删除movie表记录
// Author [yourname](https://github.com/yourname)
func (mService *MovieService) DeleteMovieByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]movie.Movie{}, "id in ?", ids).Error
	return err
}

// UpdateMovie 更新movie表记录
// Author [yourname](https://github.com/yourname)
func (mService *MovieService) UpdateMovie(m movie.Movie) (err error) {
	err = global.GVA_DB.Model(&movie.Movie{}).Where("id = ?", m.ID).Updates(&m).Error
	return err
}

// GetMovie 根据id获取movie表记录
// Author [yourname](https://github.com/yourname)
func (mService *MovieService) GetMovie(id string) (m movie.Movie, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&m).Error
	return
}

// GetMovieInfoList 分页获取movie表记录
// Author [yourname](https://github.com/yourname)
func (mService *MovieService) GetMovieInfoList(info movieReq.MovieSearch) (list []movie.Movie, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&movie.Movie{})
	var ms []movie.Movie
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&ms).Error
	return ms, total, err
}
func (mService *MovieService) GetMoviePublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
