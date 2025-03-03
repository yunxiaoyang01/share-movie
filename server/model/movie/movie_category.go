// 自动生成模板MovieCategory
package movie

// movieCategory表 结构体  MovieCategory
type MovieCategory struct {
	MovieId    *int `json:"movieId" form:"movieId" gorm:"primarykey;column:movie_id;comment:;size:19;"`          //movieId字段
	CategoryId *int `json:"categoryId" form:"categoryId" gorm:"primarykey;column:category_id;comment:;size:19;"` //categoryId字段
}

// TableName movieCategory表 MovieCategory自定义表名 movie_category
func (MovieCategory) TableName() string {
	return "movie_category"
}
