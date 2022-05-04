package db

import (
	"gorm.io/gorm"
)

type Creator interface {
	Create(ip string, port string, userName string, password string, dbName string) (*gorm.DB, error)
}

var sSelectorMap = make(map[string]Creator)

func GetCreatorByType(dbType string) Creator {
	return sSelectorMap[dbType]
}

func init() {
	sSelectorMap["mysql"] = MySQLCreator{}
}
