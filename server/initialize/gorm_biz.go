package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/movie"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(movie.Movie{}, movie.MovieCategory{})
	if err != nil {
		return err
	}
	return nil
}
