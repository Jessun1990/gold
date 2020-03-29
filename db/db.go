package db

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jessun2017/gold/logger"
	"github.com/jessun2017/gold/utils"
	"xorm.io/xorm"
)

var l *logger.Logger

func init() {
	l = logger.NewLogger("db", nil)
}

// StdConfig ...
type StdConfig struct {
	Driver   string
	Username string
	Password string
	Protocol string
	Address  string
	Port     string
	Dbname   string
	Params   string
}

// Config ...
type Config struct {
	StdConfig
	MaxIdleConns int // 最大空闲连接数
	MaxOpenConns int // 最大打开连接数
	KeepAlive    time.Duration

	// 缓存配置
	CacheType    string
	CacheMaxSize int
	CacheTimeout time.Duration

	Slaves []StdConfig
}

// InitDb ...
func InitDb() (engine xorm.EngineInterface, err error) {
	//var isNew bool

	return
}

// getEngine 根据 config 数据库的配置获取/新建要给 xorm 的 engine，用于数据库连接
func getEngine(config StdConfig) (engine *xorm.Engine, isNew bool, err error) {
	if config.Driver == "" {
		err = errors.New("db Config Driver is nil")
		return
	}

	driver := config.Driver
	dbDsn := getDBDsn(config)

	if e, ok := engineMap.Load(utils.Md5(dbDsn)); ok {
		return e.(*xorm.Engine), isNew, nil
	}

	l.Info("dbDsn: %+v", dbDsn)

	// 没有已经存在的 engine，新建一个
	engine, err = xorm.NewEngine(driver, dbDsn)
	if err != nil {
		return engine, isNew,
			fmt.Errorf("create new engine failed. [dbConfig]: %+v, [err]: %+v", config, err)
	}

	engineMap.Store(utils.Md5(dbDsn), engine)
	isNew = true

	return
}

// getDBDsn
func getDBDsn(config StdConfig) (dbDsn string) {
	config.Driver = strings.ToLower(config.Driver)
	if f, ok := dsnFuncMap[config.Driver]; ok {
		dbDsn = f(config)
		return
	}
	panic("error db driver")
}
