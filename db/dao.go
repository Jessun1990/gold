package db

import (
	"errors"
	"sync"
)

// Cond ...
type Cond map[string]interface{}

var dsnFuncMap = make(map[string]func(StdConfig) string)
var engineMap = new(sync.Map)

// ErrDaoNotInitialed ...
var ErrDaoNotInitialed = errors.New("dao not initialized")

// Dao ...
type Dao interface {
	UpdateByID(Model) (int64, error)
	UpdateByIDs(Model, Cond, []interface{}) (int64, error)
	UpdateByWhere(Model, Cond, Cond) (int64, error)

	Insert(Model) (int64, error)
	InsertMulti(interface{}) (int64, error)

	SearchOne(Model, Cond) (bool, error)
	Search(Model, interface{}, Cond) error
	SearchAndCount(Model, interface{}, Cond) (int64, error)

	GetMulti(Model, interface{}, ...interface{}) error
	Count(Model, Cond) (int64, error)

	DeleteByIDs(Model, interface{}) (int64, error)
	DeleteByWhere(Model, Cond) (int64, error)

	EnableCache(Model)
	DisableCache(Model)
	ClearCache(Model)
}
