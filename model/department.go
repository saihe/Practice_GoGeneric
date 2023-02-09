package model

type Department struct {
	Base  Base
	Phone string
}

func (*Department) TableName() string {
	return "department"
}
