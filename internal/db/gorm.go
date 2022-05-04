package db

import (
	"github.com/wujunyi792/gin-template-new/config"
	"github.com/wujunyi792/gin-template-new/internal/logger"
	"gorm.io/gorm"
)

type MainGORM struct {
	mDB *gorm.DB
}

func init() {
	if !config.GetConfig().SQL.Use {
		panic("SQL not open, please check config")
	}
}

func MustCreateGorm() *MainGORM {
	conf := config.GetConfig()
	var creator = GetCreatorByType(conf.SQL.Config.TYPE)
	if creator == nil {
		logger.Error.Fatalf("fail to find creator for type:%s", conf.SQL.Config.TYPE)
		return nil
	}
	db, err := creator.Create(conf.SQL.Config.IP, conf.SQL.Config.PORT, conf.SQL.Config.USER, conf.SQL.Config.PASSWORD, conf.SQL.Config.DATABASE)
	if err != nil {
		logger.Error.Fatalln(err)
		return nil
	}

	return &MainGORM{mDB: db}
}

func (sgo MainGORM) GetDB() *gorm.DB {
	return sgo.mDB
}
