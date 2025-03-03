
package movie

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/movie"
    movieReq "github.com/flipped-aurora/gin-vue-admin/server/model/movie/request"
)

type CategoryService struct {}
// CreateCategory 创建category表记录
// Author [yourname](https://github.com/yourname)
func (categoryService *CategoryService) CreateCategory(category *movie.Category) (err error) {
	err = global.GVA_DB.Create(category).Error
	return err
}

// DeleteCategory 删除category表记录
// Author [yourname](https://github.com/yourname)
func (categoryService *CategoryService)DeleteCategory(id string) (err error) {
	err = global.GVA_DB.Delete(&movie.Category{},"id = ?",id).Error
	return err
}

// DeleteCategoryByIds 批量删除category表记录
// Author [yourname](https://github.com/yourname)
func (categoryService *CategoryService)DeleteCategoryByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]movie.Category{},"id in ?",ids).Error
	return err
}

// UpdateCategory 更新category表记录
// Author [yourname](https://github.com/yourname)
func (categoryService *CategoryService)UpdateCategory(category movie.Category) (err error) {
	err = global.GVA_DB.Model(&movie.Category{}).Where("id = ?",category.Id).Updates(&category).Error
	return err
}

// GetCategory 根据id获取category表记录
// Author [yourname](https://github.com/yourname)
func (categoryService *CategoryService)GetCategory(id string) (category movie.Category, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&category).Error
	return
}
// GetCategoryInfoList 分页获取category表记录
// Author [yourname](https://github.com/yourname)
func (categoryService *CategoryService)GetCategoryInfoList(info movieReq.CategorySearch) (list []movie.Category, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&movie.Category{})
    var categorys []movie.Category
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&categorys).Error
	return  categorys, total, err
}
func (categoryService *CategoryService)GetCategoryPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
