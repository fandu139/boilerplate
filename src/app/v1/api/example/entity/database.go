package entity

import "time"

// Example Mapping
type Example struct {
	ID              int        `gorm:"column:id_order;primary_key" json:"id_order"`
	UUID            string     `gorm:"column:uuid;primary_key" json:"uuid"`
	ExampleNumber   string     `gorm:"column:order_number;not null;type:varchar(100)" json:"order_number"`
	UserUUID        string     `gorm:"column:user_uuid;not null;type:varchar(100)" json:"uuid_user"`
	IDExampleType   int64      `gorm:"column:id_order_type;unique;not null;type:varchar(100)" json:"id_order_type"`
	IDStatusExample int64      `gorm:"column:id_order_status;unique;not null;type:varchar(100)" json:"id_order_status"`
	CreatedAt       *time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       *time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt       *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

// ExampleStatus Mapping
type ExampleStatus struct {
	ID              int64      `gorm:"column:id_order_status;primary_key" json:"id_order_status"`
	NMStatusExample string     `gorm:"column:nm_status_order;primary_key" json:"nm_status_order"`
	CreatedAt       *time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       *time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt       *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
