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
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateRequest struct {
	Title  string `json:"title" binding:"required"`
	Detail string `json:"detail" binding:"required"`
}

type Status string

const (
	STATUS_WAIT    Status = "wait"
	STATUS_RUNNING Status = "running"
	STATUS_FINISH  Status = "finish"
)

func (s Status) CheckValid() Status {
	if (s != STATUS_WAIT) && (s != STATUS_RUNNING) && (s != STATUS_FINISH) {
		return STATUS_WAIT
	}
	return s
}
