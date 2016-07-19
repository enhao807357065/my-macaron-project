package models

type TestLh struct {
	Id                   	int         	`json:"id,omitempty"`
	Name                	string        	`json:"name" xorm:"varchar(64)"`
	Age                	int        	`json:"age,omitempty"`
}