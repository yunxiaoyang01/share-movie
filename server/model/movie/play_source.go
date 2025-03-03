
// 自动生成模板PlaySource
package movie
import (
)

// playSource表 结构体  PlaySource
type PlaySource struct {
    ID  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:19;"`  //id字段
    MovieId  *int `json:"movieId" form:"movieId" gorm:"column:movie_id;comment:影视ID;size:19;"`  //影视ID
    SourceType  *string `json:"sourceType" form:"sourceType" gorm:"column:source_type;comment:网盘类型;size:20;"`  //网盘类型
    Episode  *string `json:"episode" form:"episode" gorm:"column:episode;comment:集数/版本;size:50;"`  //集数/版本
    ResourceUrl  *string `json:"resourceUrl" form:"resourceUrl" gorm:"column:resource_url;comment:资源链接;size:500;"`  //资源链接
    ExtractCode  *string `json:"extractCode" form:"extractCode" gorm:"column:extract_code;comment:提取码;size:20;"`  //提取码
    FileSize  *int `json:"fileSize" form:"fileSize" gorm:"column:file_size;comment:文件大小(字节);size:19;"`  //文件大小(字节)
    FileFormat  *string `json:"fileFormat" form:"fileFormat" gorm:"column:file_format;comment:文件格式;size:10;"`  //文件格式
    Valid  *bool `json:"valid" form:"valid" gorm:"column:valid;comment:是否有效;"`  //是否有效
}


// TableName playSource表 PlaySource自定义表名 play_source
func (PlaySource) TableName() string {
    return "play_source"
}





