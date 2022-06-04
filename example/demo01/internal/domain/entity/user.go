package entity

import (
	"github.com/AkiOuma/demo01/internal/domain/valueobject"
	"time"
)

// Entity
type User struct {
	id        int
	name      string
	height    float64
	group     int
	createdAt time.Time
	updatedAt time.Time
}

// Constructor
func NewUser(root int, vo *valueobject.User) *User {
	return &User{
		id:        root,
		name:      vo.Name,
		height:    vo.Height,
		group:     vo.Group,
		createdAt: vo.CreatedAt,
		updatedAt: vo.UpdatedAt,
	}
}

// Getter
func (e *User) GetId() int {
	return e.id
}
func (e *User) GetName() string {
	return e.name
}
func (e *User) GetHeight() float64 {
	return e.height
}
func (e *User) GetGroup() int {
	return e.group
}
func (e *User) GetCreatedAt() time.Time {
	return e.createdAt
}
func (e *User) GetUpdatedAt() time.Time {
	return e.updatedAt
}
