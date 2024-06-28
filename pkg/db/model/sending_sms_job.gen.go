// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSendingSmsJob = "sending_sms_job"

// SendingSmsJob mapped from table <sending_sms_job>
type SendingSmsJob struct {
	ID          int32     `gorm:"column:id;primaryKey" json:"id"`
	ContactID   int32     `gorm:"column:contact_id;not null" json:"contact_id"`
	SmsQueueID  int32     `gorm:"column:sms_queue_id;not null" json:"sms_queue_id"`
	AddedAt     time.Time `gorm:"column:added_at;not null" json:"added_at"`
	SendingTime time.Time `gorm:"column:sending_time;not null" json:"sending_time"`
	Status      string    `gorm:"column:status;not null" json:"status"`
	ErrorText   string    `gorm:"column:error_text;default:NULL" json:"error_text"`
}

// TableName SendingSmsJob's table name
func (*SendingSmsJob) TableName() string {
	return TableNameSendingSmsJob
}
