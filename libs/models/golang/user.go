package models

import "time"

type UserRole string
type UserStatus string

const (
	UserRoleUser  UserRole = "user"
	UserRoleAdmin UserRole = "admin"
)

const (
	UserStatusOnline  UserStatus = "online"
	UserStatusOffline UserStatus = "offline"
)

type User struct {
	ID        uint32     `gorm:"primary_key;autoIncrement" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Email     string     `gorm:"unique" json:"email"`
	Tag       string     `gorm:"unique" json:"tag"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Role      UserRole   `gorm:"default:user" json:"role"`
	Status    UserStatus `gorm:"default:offline" json:"status"`
	LastLogin *time.Time `json:"last_login"`
}

func (User) TableName() string {
	return "users"
}

type BlockedUser struct {
	BlockedIn   uint32     `gorm:"primary_key" json:"blocked_in"`
	BlockedBy   uint32     `json:"blocked_by"`
	BlockedUser uint32     `json:"blocked_user_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`

	Chat      Chat `gorm:"foreignkey:ChatID;references:ID" json:"chat"`
	Initiator User `gorm:"foreignkey:BlockedBy;references:ID" json:"initiator"`
	User      User `gorm:"foreignkey:BlockedUser;references:ID" json:"user"`
}

func (BlockedUser) TableName() string {
	return "blocked_users"
}

type OTP struct {
	Email     string    `gorm:"primary_key" json:"email"`
	Otp       string    `json:"otp"`
	ExpiresAt time.Time `json:"expires_at"`

	User User `gorm:"foreignkey:Email;references:Email" json:"user"`
}

func (OTP) TableName() string {
	return "otps"
}
