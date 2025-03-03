package movie

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title       string       `gorm:"type:varchar(100);index;comment:影视标题"`
	Cover       string       `gorm:"type:varchar(255);comment:封面图地址"`
	PlayCount   int          `gorm:"default:0;comment:播放次数"`
	SourceType  string       `gorm:"type:varchar(20);index;comment:网盘类型"`
	ShareTime   time.Time    `gorm:"comment:分享时间"`
	ExpireTime  time.Time    `gorm:"index;comment:过期时间"`
	PlaySources []PlaySource `gorm:"foreignKey:MovieID"`
}
