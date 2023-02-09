package model

import "time"

type TableModel interface {
	TableName() string
}

type Base struct {
	ID            string
	CreatedUserID string
	UpdatedUserID string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
