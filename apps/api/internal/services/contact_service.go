package services

import (
	"context"
	"fmt"
	"log"

	"github.com/resend/resend-go/v2"
	"webifylab/api/internal/config"
	"webifylab/api/internal/models"
	"webifylab/api/internal/repositories"
	"webifylab/api/pkg/telegram"
)

type ContactService interface {
	ProcessContactSubmission(ctx context.Context, contact *models.ContactSubmission) error
}

type contactService struct {
	repo       repositories.ContactRepository
	cfg        config.Config
	telegram   *telegram.Client
	resend     *resend.Client
}

func NewContactService(repo repositories.ContactRepository, cfg config.Config) ContactService {
	var tgClient *telegram.Client
	if cfg.TelegramToken != "" && cfg.TelegramChatID != "" {
		tgClient = telegram.NewClient(cfg.TelegramToken, cfg.TelegramChatID)
	}

	var resendClient *resend.Client
	if cfg.ResendAPIKey != "" {
		resendClient = resend.NewClient(cfg.ResendAPIKey)
	}

	return &contactService{
		repo:     repo,
		cfg:      cfg,
		telegram: tgClient,
		resend:   resendClient,
	}
}

func (s *contactService) ProcessContactSubmission(ctx context.Context, contact *models.ContactSubmission) error {
	// 1. Save to database
	contact.Status = "new"
	contact.Source = "website"
	if err := s.repo.Create(ctx, contact); err != nil {
		return fmt.Errorf("failed to save contact submission: %w", err)
	}

	// 2. Send Notifications asynchronously
	go s.sendNotifications(contact)

	return nil
}

func (s *contactService) sendNotifications(contact *models.ContactSubmission) {
	// Telegram Notification
	if s.telegram != nil {
		msg := fmt.Sprintf(
			"📩 *New Contact Submission*\n\n*Name:* %s\n*Email:* %s\n*Service:* %s\n\n*Message:*\n%s",
			contact.Name, contact.Email, contact.ServiceType, contact.Message,
		)
		if err := s.telegram.SendMessage(msg); err != nil {
			log.Printf("Error sending Telegram notification: %v", err)
		}
	}

	// Email Notification via Resend
	if s.resend != nil && s.cfg.AdminEmail != "" {
		params := &resend.SendEmailRequest{
			From:    "Webifylab Notification <onboarding@resend.dev>", // Replace with verified domain in production
			To:      []string{s.cfg.AdminEmail},
			Subject: fmt.Sprintf("New Lead from Webifylab: %s", contact.Name),
			Html: fmt.Sprintf(
				"<h3>New Contact Submission</h3><p><strong>Name:</strong> %s</p><p><strong>Email:</strong> %s</p><p><strong>Service:</strong> %s</p><p><strong>Message:</strong><br>%s</p>",
				contact.Name, contact.Email, contact.ServiceType, contact.Message,
			),
			ReplyTo: contact.Email,
		}

		_, err := s.resend.Emails.Send(params)
		if err != nil {
			log.Printf("Error sending Resend email notification: %v", err)
		}
	}
}
