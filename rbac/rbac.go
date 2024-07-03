package rbac

import (
	"errors"
	rediswatcher "github.com/billcobbler/casbin-redis-watcher/v2"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

const (
	defaultDatabaseName = "casbin"
	defaultTableName    = "casbin_rule"
)

type Config struct {
	RbacPath         string `json:"rbacPath"`
	RbacChannel      string `json:"rbacChannel"`
	IsCustomCallback bool
	UpdateCallback   func(string)
	IsFiltered       bool
	Filter           []gormadapter.Filter
	Redis            redis.RedisKeyConf
	Db               *gorm.DB
}

func CustomInitRbacAndPolicy(RbacPath, RbacChannel string, db *gorm.DB, redis redis.RedisKeyConf, callback func(string2 string), filter []gormadapter.Filter) *casbin.Enforcer {
	return InitRbacAndPolicy(&Config{
		RbacPath:         RbacPath,
		RbacChannel:      RbacChannel,
		Redis:            redis,
		Db:               db,
		IsFiltered:       true,
		IsCustomCallback: true,
		UpdateCallback:   callback,
		Filter:           filter,
	})
}

func EasyInitRbacAndPolicy(RbacPath, RbacChannel string, db *gorm.DB, redis redis.RedisKeyConf) *casbin.Enforcer {
	return InitRbacAndPolicy(&Config{
		RbacPath:         RbacPath,
		RbacChannel:      RbacChannel,
		Redis:            redis,
		Db:               db,
		IsFiltered:       false,
		IsCustomCallback: false,
	})
}

func InitRbacAndPolicy(c *Config) *casbin.Enforcer {
	var err error
	rbacEnforcer, err := InitRbac(c.RbacPath, c.Db)
	watcher, err := rediswatcher.NewWatcher(c.Redis.Host, rediswatcher.Password(c.Redis.Pass), rediswatcher.Channel(c.RbacChannel))
	if err != nil {
		panic(err)
	}
	err = rbacEnforcer.SetWatcher(watcher)
	if err != nil {
		panic(err)
	}
	if c.IsCustomCallback {
		err = watcher.SetUpdateCallback(c.UpdateCallback)
		if err != nil {
			panic(err)
		}
	}
	err = rbacEnforcer.SavePolicy()
	if err != nil {
		panic(err)
	}
	if c.IsFiltered {
		err = rbacEnforcer.LoadFilteredPolicy(c.Filter)
		if err != nil {
			panic(err)
		}
	} else {
		err = rbacEnforcer.LoadPolicy()
		if err != nil {
			panic(err)
		}
	}
	return rbacEnforcer
}

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
