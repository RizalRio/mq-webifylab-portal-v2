package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"webifylab/api/internal/models"
)

type ContactRepository interface {
	Create(ctx context.Context, contact *models.ContactSubmission) error
}

type contactRepository struct {
	db *pgxpool.Pool
}

func NewContactRepository(db *pgxpool.Pool) ContactRepository {
	return &contactRepository{
		db: db,
	}
}

func (r *contactRepository) Create(ctx context.Context, contact *models.ContactSubmission) error {
	query := `
		INSERT INTO contact_submissions (
			name, email, service_type, message, source, ip_address, user_agent, referrer_url
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8
		) RETURNING id, status, created_at, updated_at
	`

	err := r.db.QueryRow(ctx, query,
		contact.Name,
		contact.Email,
		contact.ServiceType,
		contact.Message,
		contact.Source,
		contact.IPAddress,
		contact.UserAgent,
		contact.ReferrerURL,
	).Scan(
		&contact.ID,
		&contact.Status,
		&contact.CreatedAt,
		&contact.UpdatedAt,
	)

	return err
}
