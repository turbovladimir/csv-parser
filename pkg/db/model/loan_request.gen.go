// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameLoanRequest = "loan_request"

// LoanRequest mapped from table <loan_request>
type LoanRequest struct {
	ID         int32     `gorm:"column:id;primaryKey" json:"id"`
	Name       string    `gorm:"column:name;not null" json:"name"`
	Surname    string    `gorm:"column:surname;not null" json:"surname"`
	Patron     string    `gorm:"column:patron;not null" json:"patron"`
	Phone      string    `gorm:"column:phone;not null" json:"phone"`
	Birth      time.Time `gorm:"column:birth;not null" json:"birth"`
	Email      string    `gorm:"column:email;not null" json:"email"`
	AddedAt    time.Time `gorm:"column:added_at;not null" json:"added_at"`
	BirthPlace string    `gorm:"column:birth_place;not null" json:"birth_place"`
}

// TableName LoanRequest's table name
func (*LoanRequest) TableName() string {
	return TableNameLoanRequest
}
