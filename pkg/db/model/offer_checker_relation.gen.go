// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameOfferCheckerRelation = "offer_checker_relation"

// OfferCheckerRelation mapped from table <offer_checker_relation>
type OfferCheckerRelation struct {
	ID              int32     `gorm:"column:id;primaryKey" json:"id"`
	OfferID         int32     `gorm:"column:offer_id" json:"offer_id"`
	ExternalOfferID int32     `gorm:"column:external_offer_id;not null" json:"external_offer_id"`
	Checker         string    `gorm:"column:checker;not null" json:"checker"`
	AddedAt         time.Time `gorm:"column:added_at;not null" json:"added_at"`
}

// TableName OfferCheckerRelation's table name
func (*OfferCheckerRelation) TableName() string {
	return TableNameOfferCheckerRelation
}
