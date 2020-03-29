package db

import "time"

// Model ...
type Model interface {
	AutoIncrColName() string
	AutoIncrColValue() int64
	TableName() string
}

// Models ...
type Models struct {
	ID        int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	IsDeleted int       `json:"is_deleted" xorm:"not null default 0 TINYINT(1)"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP updated"`
	CreatedAt time.Time `json:"created_at" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP created"`
}

// GetID ...
func (m *Models) GetID() (val int) {
	if m == nil {
		return
	}
	return m.ID
}

// GetIsDeleted ...
func (m *Models) GetIsDeleted() (val int) {
	if m == nil {
		return
	}
	return m.IsDeleted
}

// GetUpdatedAt ...
func (m *Models) GetUpdatedAt() (val time.Time) {
	if m == nil {
		return
	}
	return m.UpdatedAt
}

// GetCreatedAt ...
func (m *Models) GetCreatedAt() (val time.Time) {
	if m == nil {
		return
	}
	return m.CreatedAt
}
