package models

import (
	"time"
)

type ContactSubmission struct {
	ID          string     `json:"id"`
	Name        string     `json:"name" validate:"required"`
	Email       string     `json:"email" validate:"required,email"`
	ServiceType string     `json:"serviceType" validate:"required"`
	Message     string     `json:"message" validate:"required"`
	Status      string     `json:"status"`
	Source      string     `json:"source"`
	IPAddress   *string    `json:"ipAddress,omitempty"`
	UserAgent   *string    `json:"userAgent,omitempty"`
	ReferrerURL *string    `json:"referrerUrl,omitempty"`
	AdminNotes  *string    `json:"adminNotes,omitempty"`
	AssignedTo  *string    `json:"assignedTo,omitempty"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty"`
}
