
// 自动生成模板Category
package movie
import (
)

// category表 结构体  Category
type Category struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`  //id字段
    Name  *string `json:"name" form:"name" gorm:"column:name;comment:分类名称;size:20;"`  //分类名称
}


// TableName category表 Category自定义表名 category
func (Category) TableName() string {
    return "category"
}





