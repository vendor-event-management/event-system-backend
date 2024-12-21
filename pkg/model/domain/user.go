package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `gorm:"type:varchar(36);primary_key" json:"id"`
	Username    string    `gorm:"type:varchar(25);unique" json:"username"`
	Password    string    `gorm:"type:varchar(100)" json:"password"`
	Role        RoleType  `gorm:"type:enum('HR', 'Vendor')" json:"role"`
	Email       string    `gorm:"type:varchar(100);unique" json:"email"`
	CompanyName string    `gorm:"type:varchar(255);" json:"company_name"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (u *User) TableName() string {
	return "users"
}
