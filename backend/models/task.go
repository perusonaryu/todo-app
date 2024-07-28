// models/user.go
package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Title     string    `json:"title"`
	Detail    string    `json:"detail"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
