package models

import (
	"errors"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	// UserStatusActive represents an active user status
	UserStatusActive = "active"
	// UserStatusInactive represents an inactive user status
	UserStatusInactive = "inactive"
	// UserStatusBanned represents a banned user status
	UserStatusBanned = "banned"

	// MinPasswordLength is the minimum required password length
	MinPasswordLength = 8
	// MaxPasswordLength is the maximum allowed password length
	MaxPasswordLength = 72
)

// Role represents user role in the system
type Role string

const (
	// RoleAdmin represents administrator role
	RoleAdmin Role = "admin"
	// RoleUser represents regular user role
	RoleUser Role = "user"
)

// User represents the user model in the system
type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"-" bson:"password"`
	Role      Role               `json:"role" bson:"role"`
	Status    string             `json:"status" bson:"status"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

// Validate checks if the user data is valid
func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("name is required")
	}

	if u.Email == "" {
		return errors.New("email is required")
	}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(u.Email) {
		return errors.New("invalid email format")
	}

	if u.Password != "" {
		if len(u.Password) < MinPasswordLength {
			return errors.New("password must be at least 8 characters long")
		}
		if len(u.Password) > MaxPasswordLength {
			return errors.New("password exceeds maximum length")
		}
	}

	if u.Role == "" {
		u.Role = RoleUser
	}

	if u.Status == "" {
		u.Status = UserStatusActive
	}

	return nil
}

// HashPassword creates a bcrypt hash of the password
func (u *User) HashPassword() error {
	if len(u.Password) == 0 {
		return errors.New("password cannot be empty")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

// ComparePassword compares the password with its hash
func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// BeforeCreate handles pre-creation tasks
func (u *User) BeforeCreate() error {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	if err := u.Validate(); err != nil {
		return err
	}

	return u.HashPassword()
}

// BeforeUpdate handles pre-update tasks
func (u *User) BeforeUpdate() error {
	u.UpdatedAt = time.Now()
	return u.Validate()
}
