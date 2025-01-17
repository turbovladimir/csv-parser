// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameContact = "contact"

// Contact mapped from table <contact>
type Contact struct {
	ID        int32     `gorm:"column:id;primaryKey" json:"id"`
	QueueID   int32     `gorm:"column:queue_id;default:NULL" json:"queue_id"`
	AddedAt   time.Time `gorm:"column:added_at;not null" json:"added_at"`
	Name      string    `gorm:"column:name;default:NULL" json:"name"`
	Phone     string    `gorm:"column:phone;not null" json:"phone"`
	ContactID string    `gorm:"column:contact_id;not null" json:"contact_id"`
	Source    string    `gorm:"column:source;default:direct" json:"source"`
}

// TableName Contact's table name
func (*Contact) TableName() string {
	return TableNameContact
}
