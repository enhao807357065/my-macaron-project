package models

type TestLh struct {
	Id                   	int         	`json:"id,omitempty"`
	Name                	string        	`json:"nm" xorm:"varchar(64)"`
	Age                	int        	`json:"ag,omitempty"`
}