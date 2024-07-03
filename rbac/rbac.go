package rbac

import (
	"errors"
	rediswatcher "github.com/billcobbler/casbin-redis-watcher/v2"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

const (
	defaultDatabaseName = "casbin"
	defaultTableName    = "casbin_rule"
)

type CasbinEngine struct {
	RbacPath         string `json:"rbacPath"`
	RbacChannel      string `json:"rbacChannel"`
	IsCustomCallback bool
	UpdateCallback   func(*casbin.Enforcer, string)
	IsFiltered       bool
	Filter           []gormadapter.Filter
	Redis            redis.RedisKeyConf
	Db               *gorm.DB
	Enforcer         *casbin.Enforcer
	Watcher          persist.Watcher
}

func NewCasbinEngine(params CasbinEngine, v ...any) *CasbinEngine {
	enforcer, err := InitRbac(params.RbacPath, params.Db, v)
	if err != nil {
		panic(err)
	}
	return &CasbinEngine{
		RbacPath:         params.RbacPath,
		RbacChannel:      params.RbacChannel,
		IsCustomCallback: params.IsCustomCallback,
		UpdateCallback:   params.UpdateCallback,
		IsFiltered:       params.IsFiltered,
		Filter:           params.Filter,
		Redis:            params.Redis,
		Db:               params.Db,
		Enforcer:         enforcer,
	}
}

func InitRbac(RbacPath string, Db *gorm.DB, v ...any) (*casbin.Enforcer, error) {
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
	adapter, err := gormadapter.NewAdapterByDBUseTableName(Db, prefix, tableName)
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer(RbacPath, adapter)
	if err != nil {
		return nil, err
	}
	err = e.LoadPolicy()
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (engine *CasbinEngine) EasyNewWatcher() *CasbinEngine {
	// note: 如果是定制的过滤器模式下，不判断是否是过滤模式，也不判断是否是定制回调模式
	engine.IsCustomCallback = false
	engine.IsFiltered = false
	return engine.NewWatcher()
}

func (engine *CasbinEngine) CustomFilterNewWatcher(filters []gormadapter.Filter) *CasbinEngine {
	// note: 如果是定制的过滤器模式下，不判断是否是过滤模式，也不判断是否是定制回调模式
	engine.IsCustomCallback = true
	engine.IsFiltered = true
	engine.Filter = filters
	engine.UpdateCallback = func(enforcer *casbin.Enforcer, msg string) {
		_ = enforcer.LoadFilteredPolicy(filters)
	}
	return engine.NewWatcher()
}

func (engine *CasbinEngine) NewWatcher() *CasbinEngine {
	watcher, err := rediswatcher.NewWatcher(engine.Redis.Host, rediswatcher.Password(engine.Redis.Pass), rediswatcher.Channel(engine.RbacChannel))
	if err != nil {
		panic(err)
	}
	err = engine.Enforcer.SetWatcher(watcher)
	if err != nil {
		panic(err)
	}
	err = engine.Enforcer.SavePolicy()
	if err != nil {
		panic(err)
	}
	if engine.IsCustomCallback {
		err = watcher.SetUpdateCallback(func(msg string) {
			engine.UpdateCallback(engine.Enforcer, msg)
		})
		if err != nil {
			panic(err)
		}
	}
	if engine.IsFiltered {
		err = engine.Enforcer.LoadFilteredPolicy(engine.Filter)
		if err != nil {
			panic(err)
		}
	} else {
		err = engine.Enforcer.LoadPolicy()
		if err != nil {
			panic(err)
		}
	}
	engine.Watcher = watcher
	return engine
}
