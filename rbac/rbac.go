package rbac

import (
	"errors"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

const (
	defaultDatabaseName = "casbin"
	defaultTableName    = "casbin_rule"
)

func InitRbac(path string, db *gorm.DB, v ...any) (*casbin.Enforcer, error) {
	prefix := ""
	tableName := ""
	if len(v) == 0 {
		prefix = ""
		tableName = defaultTableName
	} else if len(v) == 1 {
		prefix = v[0].(string)
		tableName = defaultTableName
	} else if len(v) == 2 {
		prefix = v[0].(string)
		tableName = v[1].(string)
	} else {
		return nil, errors.New("wrong parameters")
	}
	adapter, err := gormadapter.NewAdapterByDBUseTableName(db, prefix, tableName)
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer(path, adapter)
	if err != nil {
		return nil, err
	}
	err = e.LoadPolicy()
	if err != nil {
		return nil, err
	}
	return e, nil
}
