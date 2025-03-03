
package movie

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/movie"
    movieReq "github.com/flipped-aurora/gin-vue-admin/server/model/movie/request"
)

type PlaySourceService struct {}
// CreatePlaySource 创建playSource表记录
// Author [yourname](https://github.com/yourname)
func (playSourceService *PlaySourceService) CreatePlaySource(playSource *movie.PlaySource) (err error) {
	err = global.GVA_DB.Create(playSource).Error
	return err
}

// DeletePlaySource 删除playSource表记录
// Author [yourname](https://github.com/yourname)
func (playSourceService *PlaySourceService)DeletePlaySource(id string) (err error) {
	err = global.GVA_DB.Delete(&movie.PlaySource{},"id = ?",id).Error
	return err
}

// DeletePlaySourceByIds 批量删除playSource表记录
// Author [yourname](https://github.com/yourname)
func (playSourceService *PlaySourceService)DeletePlaySourceByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]movie.PlaySource{},"id in ?",ids).Error
	return err
}

// UpdatePlaySource 更新playSource表记录
// Author [yourname](https://github.com/yourname)
func (playSourceService *PlaySourceService)UpdatePlaySource(playSource movie.PlaySource) (err error) {
	err = global.GVA_DB.Model(&movie.PlaySource{}).Where("id = ?",playSource.ID).Updates(&playSource).Error
	return err
}

// GetPlaySource 根据id获取playSource表记录
// Author [yourname](https://github.com/yourname)
func (playSourceService *PlaySourceService)GetPlaySource(id string) (playSource movie.PlaySource, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&playSource).Error
	return
}
// GetPlaySourceInfoList 分页获取playSource表记录
// Author [yourname](https://github.com/yourname)
func (playSourceService *PlaySourceService)GetPlaySourceInfoList(info movieReq.PlaySourceSearch) (list []movie.PlaySource, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&movie.PlaySource{})
    var playSources []movie.PlaySource
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&playSources).Error
	return  playSources, total, err
}
func (playSourceService *PlaySourceService)GetPlaySourcePublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
