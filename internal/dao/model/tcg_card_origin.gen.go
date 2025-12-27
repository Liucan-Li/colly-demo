package model

import "time"

const TableNameTcgCardOrigin = "tcg_card_origin"

// TcgCardOrigin mapped from table <tcg_card_origin>
type TcgCardOrigin struct {
	ID            string    `gorm:"column:id;primaryKey;default:uuid()" json:"id"`
	OriginContent string    `gorm:"column:origin_content;not null" json:"origin_content"`
	CreatedAt     time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName TcgCardOrigin's table name
func (*TcgCardOrigin) TableName() string {
	return TableNameTcgCardOrigin
}
