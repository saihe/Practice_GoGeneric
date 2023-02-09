package model

type Employee struct {
	Base
	Name  string
	EMail string
}

func (*Employee) TableName() string {
	return "employee"
}
